package influxdb

import (
	"log"

	client "github.com/influxdata/influxdb1-client"
)

func NewClient(conf client.Config) (*client.Client, error) {
	con, err := client.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return con, nil
}
func (con *client.Client) Ping() (time.Duration, string, error){
	duration, version, err := con.Ping()

	return duration,version,err
}
func (con *client.Client) Query(q client.Query) (*client.Response.Results, error) {
	response, err := con.Query(q)
	if err!=nil{
		return nil,err
	}
	if response.Error() !=nil{
		return nil,response.Error()
	}
	return response.Results,nil
}
func (con *client.Client) Write(bp client.BatchPoints) (*client.Response.Results, error) {
	response, err = con.Write(bps)
	if err!=nil{
		return nil,err
	}
	if response.Error() !=nil{
		return nil,response.Error()
	}
	return response.Results,nil
}
