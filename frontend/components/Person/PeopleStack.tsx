import React, { useCallback, useEffect, useState } from "react";
import { View } from "react-native";
import { getUsers } from "../../api/users";
import Navbar from "../Layout/Navbar";
import Person from "./Person";
import { EventCardItemProps, PersonProps } from "./PersonProps";
import { getEvents } from "../../api/events";

type User = Awaited<ReturnType<typeof getUsers>>[number];
type Event = Awaited<ReturnType<typeof getEvents>>[number];

export type PeopleStackProps = {
  userId: string;
};

export default function PeopleStack({ userId }: PeopleStackProps) {
  // dummy data
  const person1: PersonProps = {
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
    promptQuestion: "What is your favorite food?",
    promptResponse:
      "I am a person \ni am from earth. I have five feet and one eye. \nIf i had one wish in this entire world it would be to eat grapes and cheese for the rest of this short life that i have left to life. ",
    interests: ["swimming", "running", "sleeping", "coding"],

    relationshipType: "Long Term Relationship",
    religion: "None",
    politicalAffiliation: "Democrat",
    alchoholFrequency: "Never",
    smokingFrequency: "Never",
    drugFrequency: "Never",
    cannabisFrequency: "Never",

    instagramUsername: "@john_doe",
    mutualEvents: [
      {
        title: "Event 1",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 2",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 1",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 2",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      }
    ],
    images: [
      {
        image:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
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
    isMatched: false,
    likesYou: false,
    handleReact: () => {}
  };

  const person2: PersonProps = {
    id: 1,
    firstName: "Jane",
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
    promptQuestion: "What is your favorite food?",
    promptResponse:
      "I am a person \ni am from earth. I have five feet and one eye. \nIf i had one wish in this entire world it would be to eat grapes and cheese for the rest of this short life that i have left to life. ",
    interests: ["swimming", "running", "sleeping", "coding"],

    relationshipType: "Long Term Relationship",
    religion: "None",
    politicalAffiliation: "Democrat",
    alchoholFrequency: "Never",
    smokingFrequency: "Never",
    drugFrequency: "Never",
    cannabisFrequency: "Never",

    instagramUsername: "@john_doe",
    mutualEvents: [
      {
        title: "Event 1",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 2",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 1",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
      {
        title: "Event 2",
        description: "This is a description",
        imageUrl:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      }
    ],
    images: [
      {
        image:
          "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
      },
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
    isMatched: false,
    likesYou: false,
    handleReact: () => {}
  };

  // use get all users endpoint
  const [people, setPeople] = useState<PersonProps[]>([person1, person2]);
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [person, setPerson] = useState<PersonProps>(people[currentCardIndex]);

  const handleReact = useCallback(
    (like: boolean) => {
      console.log("HELLO", like);
      //   const userId = people[currentCardIndex].id;
      //   const currentPersonId = userId;

      // TODO - find event swipe function
      // personSwipe(userId, currentPersonId, like).then()

      // we keep looping through people
      setCurrentCardIndex((currentCardIndex + 1) % people.length);
      setPerson(people[currentCardIndex]);
    },
    [people, currentCardIndex]
  );


  
  useEffect(() => { 
    getUsers().then((fetchedPeople: User[]) => {
        fetchedPeople.forEach((fetchedPerson, index) => {
            const events : EventCardItemProps[] = [];
            getEvents({ limit: 4, offset: index }).then((fetchedEvents: Event[]) => {
                fetchedEvents.forEach((fetchedEvent: Event) => {
                    events.push({
                        title: fetchedEvent.name,
                        description: fetchedEvent.bio,
                        imageUrl: fetchedEvent.images[0],
                    
                    })
                });
            });    
            const newPerson : PersonProps=  {
                id: fetchedPerson.id,
                firstName: fetchedPerson.firstName,
                lastName: fetchedPerson.lastName,
                age: fetchedPerson.age,
                pronouns: "they/them",
                location: "San Francisco",
                school: "UC Berkeley",
                work: "Software Engineer",
                height: {
                feet: 5,
                inches: 11
                },
                promptQuestion: "What is your favorite food?",
                promptResponse:
                "I am a person \ni am from earth. I have five feet and one eye. \nIf i had one wish in this entire world it would be to eat grapes and cheese for the rest of this short life that i have left to life. ",
                interests: ["swimming", "running", "sleeping", "coding"],
            
                relationshipType: "Long Term Relationship",
                religion: "None",
                politicalAffiliation: "Democrat",
                alchoholFrequency: "Never",
                smokingFrequency: "Never",
                drugFrequency: "Never",
                cannabisFrequency: "Never",
            
                instagramUsername: "@john_doe",
                mutualEvents: events,
                images: [
                {
                    image: fetchedPerson.images[0]  },
                {
                    image: fetchedPerson.images[1], 
                    caption: "This is a caption"
                },
                {
                    image: fetchedPerson.images[2],
                    caption: "This is a caption"
                }
                ],
                isMatched: false,
                likesYou: false,
                handleReact: () => {}
            };
            setPeople((prevPeople) => [...prevPeople, newPerson]);
        });
    })
}, [userId]);

  return (
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
        promptQuestion={person.promptQuestion}
        promptResponse={person.promptResponse}
        interests={person.interests}
        relationshipType={person.relationshipType}
        religion={person.religion}
        politicalAffiliation={person.politicalAffiliation}
        alchoholFrequency={person.alchoholFrequency}
        smokingFrequency={person.smokingFrequency}
        drugFrequency={person.drugFrequency}
        cannabisFrequency={person.cannabisFrequency}
        instagramUsername={person.instagramUsername}
        mutualEvents={person.mutualEvents}
        images={person.images}
        isMatched={person.isMatched}
        handleReact={handleReact}
        likesYou={person.likesYou}
      />

      <Navbar activePage="" />
    </View>
  );
}
