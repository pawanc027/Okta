package main

import (
    "testing"
    "github.com/bxcodec/faker/v3"
)

func TestCreateOktaUser(t *testing.T) {
    // Replace with your Okta domain and API key
    oktaDomain := "https://your-okta-domain.okta.com"
    apiKey := "your-api-key"

    client, err := NewOktaClient(oktaDomain, apiKey)
    if err != nil {
        t.Fatalf("Failed to create Okta client: %v", err)
    }

    // Generate fake user data
    firstName := faker.FirstName()
    lastName := faker.LastName()
    email := faker.Email()
    login := faker.Username()

    // Create a new Okta user
    user, err := client.CreateUser(firstName, lastName, email, login)
    if err != nil {
        t.Fatalf("Failed to create Okta user: %v", err)
    }

    // Validate the user creation
    if user.Profile.FirstName != firstName {
        t.Errorf("Expected first name %s, but got %s", firstName, user.Profile.FirstName)
    }
    if user.Profile.LastName != lastName {
        t.Errorf("Expected last name %s, but got %s", lastName, user.Profile.LastName)
    }
    if user.Profile.Email != email {
        t.Errorf("Expected email %s, but got %s", email, user.Profile.Email)
    }
    if user.Profile.Login != login {
        t.Errorf("Expected login %s, but got %s", login, user.Profile.Login)
    }
}
