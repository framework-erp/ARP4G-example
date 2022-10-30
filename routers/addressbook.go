package routers

import (
	"example/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var AddressBookService service.AddressBookService

func SetAddressBookRoutes(router *gin.RouterGroup) {
	router.GET("/addcontact", func(c *gin.Context) {
		contact, err := AddressBookService.AddContact(c, c.Query("name"), c.Query("phone"))
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contact,
		})
	})

	router.GET("/getcontact", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Query("id"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		contact, err := AddressBookService.FindContactById(c, id)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contact,
		})
	})

	router.GET("/removecontact", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Query("id"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		err = AddressBookService.RemoveContact(c, id)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	router.GET("/putcontactingroup", func(c *gin.Context) {
		contactId, err := strconv.ParseInt(c.Query("contactid"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		groupId, err := strconv.ParseInt(c.Query("groupid"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		contact, err := AddressBookService.PutContactInGroup(c, contactId, groupId)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contact,
		})
	})

	router.GET("/addgroup", func(c *gin.Context) {
		group, err := AddressBookService.AddGroup(c, c.Query("name"))
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    group,
		})
	})

	router.GET("/removegroup", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Query("id"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		err = AddressBookService.RemoveGroup(c, id)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	router.GET("/getgroups", func(c *gin.Context) {
		groups, err := AddressBookService.GetGroups(c)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    groups,
		})
	})

	router.GET("/getcontactsforgroup", func(c *gin.Context) {
		groupId, err := strconv.ParseInt(c.Query("groupid"), 10, 64)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		contacts, err := AddressBookService.GetContactsForGroup(c, groupId)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contacts,
		})
	})

	router.GET("/getfriends", func(c *gin.Context) {
		contacts, err := AddressBookService.GetContactsNotInGroup(c)
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contacts,
		})
	})

	router.GET("/querycontacts", func(c *gin.Context) {
		contacts, err := AddressBookService.QueryContacts(c, c.Query("contains"))
		if err != nil {
			c.JSON(200, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data":    contacts,
		})
	})

}
