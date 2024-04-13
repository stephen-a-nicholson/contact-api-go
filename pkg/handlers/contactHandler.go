package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

var contacts []Contact

func createContact(c *gin.Context) {
    var contact Contact
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    contact.ID = len(contacts) + 1
    contacts = append(contacts, contact)

    c.JSON(http.StatusCreated, contact)
}

func getContact(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
        return
    }

    for _, contact := range contacts {
        if contact.ID == id {
            c.JSON(http.StatusOK, contact)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}

func updateContact(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
        return
    }

    var updatedContact Contact
    if err := c.ShouldBindJSON(&updatedContact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for index, contact := range contacts {
        if contact.ID == id {
            updatedContact.ID = id
            contacts[index] = updatedContact
            c.JSON(http.StatusOK, updatedContact)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}

func deleteContact(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
        return
    }

    for index, contact := range contacts {
        if contact.ID == id {
            contacts = append(contacts[:index], contacts[index+1:]...)
            c.JSON(http.StatusNoContent, nil)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}
