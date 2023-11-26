package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	// Utiliser le mode Release pour la production
	gin.SetMode(gin.ReleaseMode)

	// Créer un routeur Gin
	r := gin.New()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://safecall-web.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Ajouter des routes
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.POST("/login", login)                   // TESTED
	r.GET("/profile/:userID", getUserProfile) // TESTED
	r.GET("/search/:userID", SearchNameEndpoint)

	r.POST("/forgetPassword", forgetPassword) // UNTESTABLE
	r.POST("/forgetPasswordCode", checkcode)  // UNTESTABLE
	r.POST("/setPassword", setPswEndpoint)
	r.POST("/editPassword", editPswEndpoint)

	r.POST("/register", register)                  // TESTED
	r.POST("/profileDescription", postDescription) // TESTED
	r.POST("/profileFullName", postFullName)       // TESTED
	r.POST("/profilePhoneNB", postPhoneNB)         // TESTED
	r.POST("/profileEmail", postEmail)             // TESTED
	r.POST("/profilePic", postProfilePic)
	r.POST("/delete", deleteUser) // TESTED

	r.POST("/manageFriend", manageFriendEndpoint) // TESTED
	r.POST("/replyFriend", replyFriendEndpoint)   // TESTED
	r.GET("/listFriends/:userID", listFriends)    // TESTED

	r.POST("/addEvent", addEventEndpoint)          // TESTED
	r.POST("/delEvent", delEventEndpoint)          // TESTED
	r.POST("/confirmEvent", confirmEvent)          // TESTED
	r.GET("/listEvent/:userID", listEventEndpoint) // TESTED

	r.POST("/AddNotification", addNotificationEndpoint) // FIXME Inform Front TESTED
	r.POST("/DelNotification", delNotificationEndpoint) // TESTED
	r.GET("/notification/:UserID", GetUserNotification) // TESTED

	r.POST("/sendMessage", PostMessage)
	r.GET("/conversations/:UserID", GetConversations)
	r.GET("/messages/:UserID/:FriendID", GetMessages)
	r.GET("/delRoom/:room", DelMessage)

	r.POST("/feedback", NewFeedback) // Tested
	r.POST("/editFeedback", EditFeedbackEndpoint)
	r.POST("/delFeedback", DelFeedback) // Tested
	r.GET("/feedback", GetFeedback)     // Tested

	r.POST("/report", NewReport)     // Tested
	r.POST("/delReport", DelReports) // Tested
	r.GET("/report", GetReports)     // Tested

	r.GET("/setupProfiler", SetupProfiler)
	r.GET("/tryCall", sendCall)

	// Configurer le serveur HTTPS
	portHTTPS := 8080
	certFile := "cert.pem"
	keyFile := "key.unencrypted.pem"

	// Configurer le serveur HTTP
	portHTTP := 7070

	var wg sync.WaitGroup

	// Lancer le serveur HTTPS dans une goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := r.RunTLS(fmt.Sprintf(":%d", portHTTPS), certFile, keyFile)
		if err != nil {
			log.Fatal("Erreur lors du démarrage du serveur HTTPS : ", err)
		}
	}()

	// Lancer le serveur HTTP dans une goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(fmt.Sprintf(":%d", portHTTP), r)
		if err != nil {
			log.Fatal("Erreur lors du démarrage du serveur HTTP : ", err)
		}
	}()

	// Attendre que les serveurs se terminent
	wg.Wait()
}

// // package main

// // import (
// // 	"net/http"

// // 	"github.com/gin-gonic/gin"
// // 	"golang.org/x/crypto/acme/autocert"
// // )

// // // This function is here for test purpose with Postman

// // func main() {
// // 	r := gin.Default()

// // 	r.Use(func(c *gin.Context) {
// // 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// // 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// // 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// // 		if c.Request.Method == "OPTIONS" {
// // 			c.AbortWithStatus(http.StatusOK)
// // 			return
// // 		}

// // 		c.Next()
// // 	})

// // 	// Chemin vers les fichiers de certificat et de clé privée
// // 	certFile := "cert.pem"
// // 	keyFile := "key.unencrypted.pem"
// // 	// Password : safecall

// // 	// Utilisez autocert pour la gestion automatique des certificats (let's encrypt)
// // 	// En production, remplacez le domaine factice par votre propre domaine
// // 	m := autocert.Manager{
// // 		Prompt:     autocert.AcceptTOS,
// // 		HostPolicy: autocert.HostWhitelist("*"),
// // 		Cache:      autocert.DirCache("certs"), // Emplacement pour stocker les certificats
// // 	}

// // 	// Lancez le serveur avec la gestion automatique des certificats
// // 	go http.ListenAndServe(":80", m.HTTPHandler(nil))

// // 	r.POST("/login", login)                   // TESTED
// // 	r.GET("/profile/:userID", getUserProfile) // TESTED
// // 	r.GET("/search/:userID", SearchNameEndpoint)

// // 	r.POST("/forgetPassword", forgetPassword) // UNTESTABLE
// // 	r.POST("/forgetPasswordCode", checkcode)  // UNTESTABLE
// // 	r.POST("/setPassword", setPswEndpoint)
// // 	r.POST("/editPassword", editPswEndpoint)

// // 	r.POST("/register", register)                  // TESTED
// // 	r.POST("/profileDescription", postDescription) // TESTED
// // 	r.POST("/profileFullName", postFullName)       // TESTED
// // 	r.POST("/profilePhoneNB", postPhoneNB)         // TESTED
// // 	r.POST("/profileEmail", postEmail)             // TESTED
// // 	r.POST("/profilePic", postProfilePic)
// // 	r.POST("/delete", deleteUser) // TESTED

// // 	r.POST("/manageFriend", manageFriendEndpoint) // TESTED
// // 	r.POST("/replyFriend", replyFriendEndpoint)   // TESTED
// // 	r.GET("/listFriends/:userID", listFriends)    // TESTED

// // 	r.POST("/addEvent", addEventEndpoint)          // TESTED
// // 	r.POST("/delEvent", delEventEndpoint)          // TESTED
// // 	r.POST("/confirmEvent", confirmEvent)          // TESTED
// // 	r.GET("/listEvent/:userID", listEventEndpoint) // TESTED

// // 	r.POST("/AddNotification", addNotificationEndpoint) // FIXME Inform Front TESTED
// // 	r.POST("/DelNotification", delNotificationEndpoint) // TESTED
// // 	r.GET("/notification/:UserID", GetUserNotification) // TESTED

// // 	r.POST("/sendMessage", PostMessage)
// // 	r.GET("/conversations/:UserID", GetConversations)
// // 	r.GET("/messages/:UserID/:FriendID", GetMessages)
// // 	r.GET("/delRoom/:room", DelMessage)

// // 	r.POST("/feedback", NewFeedback) // Tested
// // 	r.POST("/editFeedback", EditFeedbackEndpoint)
// // 	r.POST("/delFeedback", DelFeedback) // Tested
// // 	r.GET("/feedback", GetFeedback)     // Tested

// // 	r.POST("/report", NewReport)     // Tested
// // 	r.POST("/delReport", DelReports) // Tested
// // 	r.GET("/report", GetReports)     // Tested

// // 	r.GET("/setupProfiler", SetupProfiler)
// // 	r.GET("/tryCall", sendCall)

// // 	r.RunTLS(":8080", certFile, keyFile)
// // }

// // // http.ListenAndServeTLS(":443", certFile, keyFile, r)
