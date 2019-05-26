package repository

import (
	"demoGo/configuration"
	"fmt"
	"time"
)

//ClientConfig represents client config
type ClientConfig struct {
	LinkedToSpring bool `json:"linked_to_spring"`
}

//Client struct it implements the osin client
type Client struct {
	ID               int64        `xorm:"'id' pk autoincr"`
	ClientID         string       `xorm:"'client_id'" json:"clientId"`
	Name             string       `xorm:"'name'" json:"name"`
	SpringClientName string       `xorm:"'spring_client_name'" json:"spring_client_name"`
	ClientSecret     string       `xorm:"'client_secret'" json:"clientSecret"`
	ILPSecret        string       `xorm:"'ilp_secret'" json:"ilpSecret"`
	OrganizationID   string       `xorm:"'organization_id'" json:"organizationID"`
	RedirectURL      string       `xorm:"'redirect_url'" json:"redirectUrl"`
	IsActive         bool         `xorm:"'is_active'" json:"isActive"`
	Meta             ClientConfig `xorm:"json" json:"meta"`
	CreatedAt        time.Time    `xorm:"'created_at' created" json:"createdAt"`
	UpdatedAt        time.Time    `xorm:"'updated_at' updated" json:"updatedAt"`
}

const (
	clientTableName = "client"
)

func ClientIndex() []*Client {
	var clients []*Client
	sess := configuration.GetConnection()
	defer sess.Close()
	sess.Table(clientTableName)
	err := sess.Find(&clients)
	if err != nil {
		fmt.Println(err)
	}
	return clients
}
