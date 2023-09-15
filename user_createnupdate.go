package main

import (
    "testing"
    "github.com/bxcodec/faker/v3"
)

func TestCreateAndUpdateOktaUser(t *testing.T) {
    // Replace with your Okta domain and API key
    oktaDomain := "https://your-okta-domain.okta.com"
    apiKey := "your-api-key"

    client, err := NewOktaClient(oktaDomain, apiKey)
    if err != nil {
        t.Fatalf("Failed to create Okta client: %v", err)
    }

    // Generate fake user data for creation
    firstNameCreate := faker.FirstName()
    lastNameCreate := faker.LastName()
    emailCreate := faker.Email()
    loginCreate := faker.Username()

    // Create a new Okta user
    user, err := client.CreateUser(firstNameCreate, lastNameCreate, emailCreate, loginCreate)
    if err != nil {
        t.Fatalf("Failed to create Okta user: %v", err)
    }

    // Validate the user creation
    if user.Profile.FirstName != firstNameCreate {
        t.Errorf("Expected first name %s, but got %s", firstNameCreate, user.Profile.FirstName)
    }
    if user.Profile.LastName != lastNameCreate {
        t.Errorf("Expected last name %s, but got %s", lastNameCreate, user.Profile.LastName)
    }
    if user.Profile.Email != emailCreate {
        t.Errorf("Expected email %s, but got %s", emailCreate, user.Profile.Email)
    }
    if user.Profile.Login != loginCreate {
        t.Errorf("Expected login %s, but got %s", loginCreate, user.Profile.Login)
    }

    // Generate fake user data for update
    firstNameUpdate := faker.FirstName()
    lastNameUpdate := faker.LastName()

    // Update the user's first and last name
    updatedUser, err := client.UpdateUserFirstNameAndLastName(user.Id, firstNameUpdate, lastNameUpdate)
    if err != nil {
        t.Fatalf("Failed to update Okta user: %v", err)
    }

    // Validate the user update
    if updatedUser.Profile.FirstName != firstNameUpdate {
        t.Errorf("Expected updated first name %s, but got %s", firstNameUpdate, updatedUser.Profile.FirstName)
    }
    if updatedUser.Profile.LastName != lastNameUpdate {
        t.Errorf("Expected updated last name %s, but got %s", lastNameUpdate, updatedUser.Profile.LastName)
    }
}
