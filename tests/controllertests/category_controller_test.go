package controllertests

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/victorsteven/fullstack/api/auth"
)

func TestGetAllCategory(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://localhost:8080/categories")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	type Category struct {
		Id        uint64 `json:"id"`
		Category  string `json:"category"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var result []Category
	// result := []map[string]interface{}{}
	if err = json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	// t.Log(string(body))
	// t.Log(result)
	// t.Logf("%+v", result)

	// Test case
	if resp.StatusCode != 200 {
		t.Fatal(resp.StatusCode)
	}

	if len(result) != 3 {
		t.Fatal(len(result))
	}

}

func TestGetAllCategoryWithEmptyResult(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://localhost:8080/categories")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var result []string
	if err = json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(result)

	// Test case
	if resp.StatusCode != 200 {
		t.Fatal(resp.StatusCode)
	}

	if len(result) > 0 {
		t.Fatal(len(result))
	}
}

func TestGetCategoryWithValue(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	input := strconv.Itoa(1)

	resp, err := http.Get("http://localhost:8080/categories/" + input)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := map[string]interface{}{}
	if err = json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	// t.Log(result)
	// t.Log(result["id"])
	// t.Logf("%T", result["id"])
	// t.Log(result["category"])
	// t.Log(result["created_at"])
	// t.Log(result["updated_at"])

	// Test case
	if resp.StatusCode != 200 {
		t.Fatal(resp.StatusCode)
	}

	id, err := strconv.ParseFloat(input, 64)
	if err != nil {
		t.Fatal(err)
	}

	if result["id"] != id {
		t.Fatal(result["id"])
	}

	if result["category"] != "Hello" {
		t.Fatal(result["category"])
	}

}

func TestGetCategoryWithValueButTheDataIsNotFound(t *testing.T) {
	var err error
	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}
	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	input := strconv.Itoa(4)

	resp, err := http.Get("http://localhost:8080/categories/" + input)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}

}

func TestGetCategoryWithValueButTheValueIsAlphabetic(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	input := "a"

	resp, err := http.Get("http://localhost:8080/categories/" + input)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}

}

func TestGetCategoryWithValueButTheValueIsSymbol(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	input := "-"

	resp, err := http.Get("http://localhost:8080/categories/" + input)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}

}

func TestCreateCategoryWithAcceptedJsonFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case

	input := map[string]interface{}{
		"category": "Yup",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	resp, err := http.Post("http://localhost:8080/categories", "application/json", read)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	// t.Log(string(body))

	// Test case
	if resp.StatusCode != 201 {
		t.Fatal(resp.StatusCode)
	}

	if result["category"] != input["category"] {
		t.Fatal(result["category"])
	}

}

func TestCreateCategoryWithUnacceptableJsonFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case

	input := map[string]interface{}{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	resp, err := http.Post("http://localhost:8080/categories", "application/json", read)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}

}

func TestCreateCategoryWithOtherFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	input := "<input><category>Yup</category></input>"

	read := strings.NewReader(input)

	resp, err := http.Post("http://localhost:8080/categories", "text/xml", read)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}

}

func TestCreateCategoryWithEmptyCategoryInput(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case

	input := map[string]interface{}{
		"category": "",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	resp, err := http.Post("http://localhost:8080/categories", "application/json", read)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}

}

func TestCreateCategoryWithNumericCategoryInput(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Input case

	input := map[string]interface{}{
		"category": 14,
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	resp, err := http.Post("http://localhost:8080/categories", "application/json", read)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}

}

func TestUpdateCategoryWithAcceptableJsonFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(2)

	input := map[string]interface{}{
		"category": "Dunia",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	// t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 204 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithUnacceptableJsonFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(2)

	input := map[string]interface{}{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithOtherFormat(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(2)

	input := "<input><category>Yup</category></input>"

	read := strings.NewReader(input)

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithEmptyCategoryInput(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(2)

	input := map[string]interface{}{
		"category": "",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithNumericCategoryInput(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(2)

	input := map[string]interface{}{
		"category": 1,
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 422 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithEmptyIdValue(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := ""

	input := map[string]interface{}{
		"category": "Dunia",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 404 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithIdValueButTheValueIsAlphabetic(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := "a"

	input := map[string]interface{}{
		"category": "Dunia",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}
}

func TestUpdateCategoryWithIdValueButTheValueIsSymbol(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := ":"

	input := map[string]interface{}{
		"category": "Dunia",
	}

	format, err := json.Marshal(input)

	read := strings.NewReader(string(format))

	req, err := http.NewRequest("PUT", "http://localhost:8080/categories/"+id, read)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}
}

func TestDeleteCategoryWithIdValue(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := strconv.Itoa(1)

	req, err := http.NewRequest("DELETE", "http://localhost:8080/categories/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Uncomment the code below to see the result
	t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 204 {
		t.Fatal(resp.StatusCode)
	}
}

func TestDeleteCategoryWithIdValueButTheValueIsAlphabetic(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := "a"

	req, err := http.NewRequest("DELETE", "http://localhost:8080/categories/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}
}

func TestDeleteCategoryWithIdValueButTheValueIsSymbol(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// jwt key for authorization
	key, err := auth.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	// Input case
	id := ":"

	req, err := http.NewRequest("DELETE", "http://localhost:8080/categories/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+key)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
	}

	// Uncomment the code below to see the result
	t.Log(string(body))
	// t.Log(resp.StatusCode)

	// Test case
	if resp.StatusCode != 400 {
		t.Fatal(resp.StatusCode)
	}
}
