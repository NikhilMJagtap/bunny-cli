package client

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/fatih/color"
    "github.com/rodaine/table"
    "github.com/spf13/cobra"
)

type BunnyClient struct {
    host    string
    headers map[string]string
}

var bunnyClient *BunnyClient = nil

func (b BunnyClient) unmarshalResponse(resp *http.Response) (interface{}, error) {
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    body = bytes.TrimLeft(body, " \t\r\n")
    if err != nil {
        return nil, err
    }
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

func (b BunnyClient) Get(path string, queryParams map[string]interface{}) (interface{}, error) {
    c := http.Client{}
    req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", b.host, path), nil)
    for header, val := range b.headers {
        req.Header.Add(header, val)
    }
    q := req.URL.Query()
    if queryParams != nil && len(queryParams) > 0 {
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

func GetBunnyClient() *BunnyClient {
    if bunnyClient == nil {
        bunnyClient = &BunnyClient{
            "https://api.bunny.net",
            map[string]string{
                "AccessKey": "",
                "Accept":    "application/json",
            },
        }
    }
    return bunnyClient
}
