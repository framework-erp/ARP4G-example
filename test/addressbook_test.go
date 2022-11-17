package test

import (
	"context"
	"example/infrastructure/arpservice"
	"example/service"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/framework-arp/ARP4G/test"
)

func TestPlaceOrder(t *testing.T) {
	node, err := snowflake.NewNode(1)
	var addressBookService service.AddressBookService
	addressBookService = arpservice.NewArpAddressBookService(nil, node)

	contact, err := addressBookService.AddContact(context.Background(), "neo", "12345")
	test.AssertNoError(t, err)

	group, err := addressBookService.AddGroup(context.Background(), "friend")
	test.AssertNoError(t, err)

	contact, err = addressBookService.PutContactInGroup(context.Background(), contact.Id, group.Id)
	test.AssertNoError(t, err)
	test.AssertEqual(t, group.Id, contact.GroupId)

	err = addressBookService.RemoveGroup(context.Background(), group.Id)
	test.AssertNoError(t, err)

	group, err = addressBookService.AddGroup(context.Background(), "family")
	test.AssertNoError(t, err)

	contact, err = addressBookService.PutContactInGroup(context.Background(), contact.Id, group.Id)
	test.AssertNoError(t, err)
	test.AssertEqual(t, group.Id, contact.GroupId)
}
