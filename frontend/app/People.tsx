import React from "react";
import { View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import Person from "../components/Person/Person";
import { PersonProps } from "../components/Person/PersonProps";

export default function People() {
  const person: PersonProps = {
    id: 1,
    firstName: "John",
    lastName: "Doe",
    age: 25,
    pronouns: "he/him",
    location: "San Francisco",
    school: "UC Berkeley",
    work: "Software Engineer",
    height: {
      feet: 5,
      inches: 11
    },
    bio: "I am a person \ni am from earth. I have five feet and one eye. \nIf i had one wish in this entire world it would be to eat grapes and cheese for the rest of this short life that i have left to life. ",
    interests: ["swimming", "running", "sleeping", "coding"],

    relationshipType: "Long Term Relationship",
    religion: "None",
    politicalAffiliation: "Democrat",
    alchoholFrequency: "Never",
    smokingFrequency: "Never",
    drugFrequency: "Never",
    marijuanaFrequency: "Never",

    instagram: "instagram.com",
    mutualEvents: [],
    images: [
      {
        image:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
        caption: "This is a caption"
      },
      {
        image:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
        caption: "This is a caption"
      }
    ], 
    isMatched: false
  };
  return (
    // foreach ... add or...
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Person
        id={person.id}
        firstName={person.firstName}
        lastName={person.lastName}
        age={person.age}
        pronouns={person.pronouns}
        location={person.location || ""}
        school={person.school || ""}
        work={person.work || ""}
        height={person.height || undefined}
        bio={person.bio}
        interests={person.interests}
        relationshipType={person.relationshipType}
        religion={person.religion}
        politicalAffiliation={person.politicalAffiliation}
        alchoholFrequency={person.alchoholFrequency}
        smokingFrequency={person.smokingFrequency}
        drugFrequency={person.drugFrequency}
        marijuanaFrequency={person.marijuanaFrequency}
        instagram={person.instagram}
        mutualEvents={person.mutualEvents}
        images={person.images}
        isMatched={person.isMatched}
      />
      {/* <Text style={{ fontFamily: "DMSansRegular" }}>People</Text> */}

      <Navbar activePage="" />
    </View>
  );
}
