import React, { useCallback, useEffect, useState } from "react";
import { ActivityIndicator, View } from "react-native";
import { getEvents } from "../../api/events";
import { getUsers } from "../../api/users";
import Navbar from "../Layout/Navbar";
import Person from "./Person";
import { EventCardItemProps, PersonProps } from "./PersonProps";

type User = Awaited<ReturnType<typeof getUsers>>[number];
type Event = Awaited<ReturnType<typeof getEvents>>[number];

export type PeopleStackProps = {
  userId: string;
};

export default function PeopleStack({ userId }: PeopleStackProps) {


  const [people, setPeople] = useState<PersonProps[]>([]);
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [person, setPerson] = useState<PersonProps>(people[currentCardIndex]);
  const [isLoading, setIsLoading] = useState(true);

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
        const events: EventCardItemProps[] = [];
        getEvents({ limit: 4, offset: index }).then((fetchedEvents: Event[]) => {
          fetchedEvents.forEach((fetchedEvent: Event) => {
            events.push({
              title: fetchedEvent.name,
              description: fetchedEvent.bio,
              imageUrl: fetchedEvent.images[0]
            });
          });
        });
        const newPerson: PersonProps = {
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
              image: fetchedPerson.images[0]
            },
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
    });
  }, [userId]);

  useEffect(() => {
    if (people.length > 0) {
      setPerson(people[currentCardIndex]);
      setIsLoading(false);
    }
  }, [people, currentCardIndex, person]);

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

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
