package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"example.com/greetings/models"

	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

/* func TestGetBlogs(t *testing.T) {
	Convey("Get Blogs", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		blog1 := models.Blog{
			ID:          GenerateUUID(8),
			Title:       "baslik1",
			Description: "description1",
		}

		blog2 := models.Blog{
			ID:          GenerateUUID(8),
			Title:       "baslik2",
			Description: "description2",
		}

		repository.CreateBlog(blog1)
		repository.CreateBlog(blog2)

		Convey("When the get request sent ", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/blogs", nil)

			resp, err := app.Test(req)
			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})
			Convey("Then the reques should return all blogs", func() {
				actualResult := []models.Blog{}
				actualResponseBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualResponseBody, &actualResult)
				So(err, ShouldBeNil)

				So(actualResult, ShouldHaveLength, 2)
				So(actualResult[0].ID, ShouldEqual, blog1.ID)
				So(actualResult[0].Title, ShouldEqual, blog1.Title)
				So(actualResult[0].Description, ShouldEqual, blog1.Description)
				So(actualResult[1].ID, ShouldEqual, blog2.ID)
				So(actualResult[1].Title, ShouldEqual, blog2.Title)
				So(actualResult[1].Description, ShouldEqual, blog2.Description)

			})

		})

	})

}

// Tekil Blog Get Etme Method Get
func TestGetBlog(t *testing.T) {
	Convey("Get blog", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		blog1 := models.Blog{
			ID:          GenerateUUID(8),
			Title:       "blog1",
			Description: "blog1 deneme",
		}
		repository.CreateBlog(blog1)

		Convey("When the get request sent ", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/blog/"+blog1.ID, nil)
			resp, err := app.Test(req, 30000)

			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product should be returned", func() {
				actualResult := models.Blog{}
				actualRespBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualRespBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult.ID, ShouldEqual, blog1.ID)
				So(actualResult.Title, ShouldEqual, blog1.Title)
				So(actualResult.Description, ShouldEqual, blog1.Description)
			})
		})
	})

}

// Tekil Blog Update Etme Method Put
func TestUpdateBlog(t *testing.T) {
	Convey("Update blog", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		blog1 := models.Blog{
			ID:          GenerateUUID(8),
			Title:       "blog1",
			Description: "blog1 deneme",
		}
		repository.CreateBlog(blog1)

		Convey("when the put request sent", func() {
			app := SetupApp(&api)

			blog2 := models.BlogDTO{
				Title:       "blog2",
				Description: "blog2 deneme",
			}
			reqBody, err := json.Marshal(blog2)

			So(err, ShouldBeNil)
			req, _ := http.NewRequest(http.MethodPut, "/blogs/"+blog1.ID, bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")

			resp, err := app.Test(req, 30000)

			So(err, ShouldBeNil)

			Convey("then status should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})
			Convey("Then product should be updated", func() {
				actualResult := models.Blog{}

				respBody, _ := ioutil.ReadAll(resp.Body)

				err = json.Unmarshal(respBody, &actualResult)
				So(err, ShouldBeNil)
				So(actualResult.ID, ShouldEqual, blog1.ID)
				So(actualResult.Title, ShouldEqual, "blog2")
				So(actualResult.Description, ShouldEqual, "blog2 deneme")

			})
		})
	})
}
*/
// Tekil Blog Delete Etme Method Delete
/* func TestDeleteBlog(t *testing.T) {
	Convey("Delete blog that user wants", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		blog1 := models.Blog{
			ID:          GenerateUUID(8),
			Title:       "blog1",
			Description: "blog1 deneme",
		}

		repository.CreateBlog(blog1)

		Convey("When the delete request sent ", func() {
			app := SetupApp(&api)

			req, _ := http.NewRequest(http.MethodDelete, "/blogs/"+blog1.ID, nil)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 204", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusNoContent)
			})

			Convey("Then phone should be deleted", func() {
				blogs, _ := repository.GetBlogs()
				So(blogs, ShouldHaveLength, 0)
				So(blogs, ShouldResemble, []models.Blog{})
			})
		})
	})
}
*/
// Tekil Blog Ekleme Etme Method Post
func TestAddUser(t *testing.T) {
	Convey("Add user", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		user1 := models.User{
			ID:       GenerateUUID(8),
			Name:     "name",
			Surname:  "SURNAME",
			Email:    "email",
			Tel:      "123",
			Password: "password",
		}

		Convey("when the post request send", func() {

			reqBody, err := json.Marshal(user1)

			req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

			app := SetupApp(&api)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 201", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusCreated)
			})

			Convey("Then Added blog should return", func() {
				actualResult, err := repository.GetByUserId(user1.ID)

				So(err, ShouldBeNil)
				So(actualResult, ShouldHaveLength, 1)
				So(actualResult, ShouldNotBeNil)
				So(actualResult.ID, ShouldEqual, user1.ID)
				So(actualResult.Name, ShouldEqual, user1.Name)
				So(actualResult.Surname, ShouldEqual, user1.Surname)
				So(actualResult.Password, ShouldEqual, user1.Password)
				So(actualResult.Tel, ShouldEqual, user1.Tel)
				So(actualResult.Email, ShouldEqual, user1.Email)
			})
		})
	})
}
