package main

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float64 `json:"price"`
}

// albums map to seed record album data.
var albums = map[string]album{
        "1": {ID: "1", Title: "norrm", Artist: "norrm Coltrane", Price: 69.69},
        "2": {ID: "2", Title: "atrash", Artist: "atrash Mulligan", Price: 17.99},
        "3": {ID: "3", Title: "Aghsham", Artist: "Aghsham Vaughan", Price: 99.99},
}

func main() {
        router := gin.Default()
        router.GET("/albums", getAlbums)
        router.GET("/albums/:id", getAlbumByID)
        router.POST("/albums", postAlbums)
        router.DELETE("/albums/:id", deleteAlbumByID)
        router.PUT("/albums/:id", updateAlbumByID)

        router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
        var newAlbum album

        // Call BindJSON to bind the received JSON to
        // newAlbum.
        if err := c.BindJSON(&newAlbum); err != nil {
                c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
                return
        }

        // Add the new album to the slice.
        albums[newAlbum.ID] = newAlbum
        c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Retrieve the album whose ID value matches the parameter.
        a, found := albums[id]
        if !found {
                c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
                return
        }

        c.IndentedJSON(http.StatusOK, a)
}

// deleteAlbumByID deletes an album whose ID value matches the id
// parameter sent by the client. If no matching album is found, an
// error is returned.
func deleteAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Delete the album whose ID value matches the parameter.
        _, found := albums[id]
        if !found {
                c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
                return
        }

        delete(albums, id)
        c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}

// updateAlbumByID updates an album whose ID value matches the id
// parameter sent by the client.
func updateAlbumByID(c *gin.Context) {
        id := c.Param("id")

        // Retrieve the existing album struct matching the id.
        a, found := albums[id]
        if !found {
                c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
                return
        }

        // Update the existing album with the new data sent from the client.
        if err := c.BindJSON(&a); err != nil {
                c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
                return
        }

        albums[id] = a
        c.IndentedJSON(http.StatusOK, gin.H{"message": "album Updated"})
}

