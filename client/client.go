package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type BunnyClient struct {
	host    string
	headers map[string]string
}

var bunnyClient *BunnyClient = nil

// A response from the BunnyCDN API is unmarshalled into either
//   - a map[string]interface{} or
//   - a []interface{}
//
// depending on the response type.
// The response is returned as an interface{} to allow for flexibility in the return type.
// In case of an error, nil is returned along with the error.
func (b BunnyClient) unmarshalResponse(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	body = bytes.TrimLeft(body, " \t\r\n")
	if err != nil {
		return nil, err
	}
	// Check if the response is an array.
	if body[0] == '[' {
		var list []interface{}
		if err := json.Unmarshal(body, &list); err != nil {
			return nil, err
		}
		return list, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// Sends a GET request to the BunnyCDN API.
// The `path` is the relative path to the BunnyCDN API.
// The `queryParams` is a map of query parameters to be added to the request.
func (b BunnyClient) Get(path string, queryParams map[string]interface{}) (interface{}, error) {
	c := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", b.host, path), nil)
	if err != nil {
		return nil, err
	}
	for header, val := range b.headers {
		req.Header.Add(header, val)
	}
	q := req.URL.Query()
	if len(queryParams) > 0 {
		for query, val := range queryParams {
			q.Add(query, fmt.Sprintf("%v", val))
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return b.unmarshalResponse(resp)
}

// Sends a POST request to the BunnyCDN API.
// The `path` is the relative path to the BunnyCDN API.
// The `data` is the data to be sent in the request body.
func (b BunnyClient) Post(path string, data interface{}) (interface{}, error) {
	c := http.Client{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", b.host, path), strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}
	for header, val := range b.headers {
		req.Header.Add(header, val)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 204 {
		return nil, nil
	}
	return b.unmarshalResponse(resp)
}

// Sends a DELETE request to the BunnyCDN API.
// The `path` is the relative path to the BunnyCDN API.
// The `data` is the data to be sent in the request body.
func (b BunnyClient) Delete(path string, data interface{}) (interface{}, error) {
	c := http.Client{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", b.host, path), strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}
	for header, val := range b.headers {
		req.Header.Add(header, val)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 204 {
		return nil, nil
	}
	return b.unmarshalResponse(resp)
}

// Handles the output of the command.
// The `cmd` is the command that is being executed.
// The `output` is the output of the command.
// The `columns` is the columns to be printed in the table.
func (b BunnyClient) HandleCommandOutput(cmd *cobra.Command, output interface{}, columns []string) error {
	isTable, err := cmd.Flags().GetBool("table")
	if err != nil {
		return err
	}
	if isTable {
		switch output.(type) {
		case []interface{}:
			var tableData []map[string]interface{}
			arrayOutput := output.([]interface{})
			for _, item := range arrayOutput {
				tableData = append(tableData, item.(map[string]interface{}))
			}
			b.PrintTable(columns, tableData)
		case map[string]interface{}:
			b.PrintTable(columns, []map[string]interface{}{output.(map[string]interface{})})
		case string:
			cmd.Println(output.(string))
		default:
			break
		}
	} else {
		jsonString, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}
		cmd.Println(string(jsonString))
	}
	return nil
}

// Utility to print the data in tabular format.
// The required columns should be provided in the columnNames array.
// A map of data (values) should be provided with the name of the column as the key and the value as the value.
func (b BunnyClient) PrintTable(columnNames []string, values []map[string]interface{}) {
	interfaceCols := make([]interface{}, len(columnNames))
	for i, v := range columnNames {
		interfaceCols[i] = v
	}

	tbl := table.New(interfaceCols...)
	headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt)

	for _, value := range values {
		var rowValues []interface{}
		for _, column := range columnNames {
			switch column {
			case "Id":
				value[column] = fmt.Sprintf("%.0f", value[column])
			}
			rowValues = append(rowValues, value[column])
		}
		tbl.AddRow(rowValues...)
	}
	tbl.Print()
}

// Returns a singleton instance of the BunnyClient.
// The API key is expected to be set in the BUNNY_NET_API_KEY environment variable.
func GetBunnyClient() *BunnyClient {
	if bunnyClient == nil {
		bunnyClient = &BunnyClient{
			"https://api.bunny.net",
			map[string]string{
				"AccessKey": os.Getenv("BUNNY_NET_API_KEY"),
				"Accept":    "application/json",
			},
		}
	}
	return bunnyClient
}
