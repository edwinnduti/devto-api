/*
Author :Edwin Nduti
Date   : Aug 2020

Description :
	 A dev.to consuming API code.
*/

package main

//imports
import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
	"log"
)

//constants
const (
	baseURL string = "https://dev.to/api/articles?"
	Username string = "username=nduti"
)

func main(){

	//instantiate Client
	client := &http.Client{}

	//formulate full URL
	url := baseURL+Username

	//Make a request
	req,err := http.NewRequest("GET",url,nil)
	Check(err)

	//get response
	resp,err := client.Do(req)
	Check(err)

	//close connection
	defer resp.Body.Close()

	var target interface{}

	//decode response to target interface
	json.NewDecoder(resp.Body).Decode(&target)

	//display Json
	fmt.Println(target)

}

//err handler
func Check(e error){
	if e!= nil{
		log.Fatalln(e)
	}
}


// Articles is a list of articles.
type Articles []struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                []string     `json:"tag_list"`
	Tags                   string       `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// Article just an article
type Article struct {
	TypeOf                 string       `json:"type_of"`
	ID                     int          `json:"id"`
	Title                  string       `json:"title"`
	Description            string       `json:"description"`
	CoverImage             string       `json:"cover_image"`
	Published              bool         `json:"published"`
	PublishedAt            time.Time    `json:"published_at"`
	TagList                string       `json:"tag_list"`
	Tags                   []string     `json:"tags"`
	Slug                   string       `json:"slug"`
	Path                   string       `json:"path"`
	URL                    string       `json:"url"`
	CanonicalURL           string       `json:"canonical_url"`
	CommentsCount          int          `json:"comments_count"`
	PositiveReactionsCount int          `json:"positive_reactions_count"`
	PublishedTimestamp     time.Time    `json:"published_timestamp"`
	User                   User         `json:"user"`
	Organization           Organization `json:"organization,omitempty"`
	FlareTag               FlareTag     `json:"flare_tag,omitempty"`
}

// User struct contains information about user
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`    // 640x640
	ProfileImage90  string `json:"profile_image_90"` // 90x90
}

// Organization struct contains information about organization (may be empty)
type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`    // 640x640
	ProfileImage90 string `json:"profile_image_90"` // 90x90
}

// FlareTag struct contains information about flare tag (may be empty)
type FlareTag struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

