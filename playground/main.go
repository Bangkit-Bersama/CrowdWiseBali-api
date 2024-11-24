package main

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func main() {
	// Replace with your actual Google Maps API key
	apiKey := "AIzaSyBBKGZvJjybn03xeoFJ-kNJi2CR0sU30vc"

	// Initialize Google Maps client
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Google Maps client: %v", err)
	}

	// Define the location (latitude, longitude) and radius for the search
	location := &maps.LatLng{
		Lat: -8.506939, // Example: Ubud, Bali
		Lng: 115.262482,
	}
	radius := 5000 // Search within 5 km

	// Call the Nearby Search API
	req := &maps.NearbySearchRequest{
		Location: location,
		Radius:   uint(radius),
		Type:     "restaurant", // Change type as needed (e.g., "hotel", "tourist_attraction")
	}
	response, err := client.NearbySearch(context.Background(), req)
	if err != nil {
		log.Fatalf("Nearby Search request failed: %v", err)
	}

	// Print the results
	log.Println("Nearby Places:")
	for _, result := range response.Results {
		log.Printf("Name: %s, Address: %s, Rating: %.1f",
			result.Name, result.Vicinity, result.Rating)
	}
}
