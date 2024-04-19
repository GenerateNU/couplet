import React from "react";
import { StyleSheet, Text, View } from "react-native";

import { router } from "expo-router";
import scaleStyleSheet from "../../scaleStyles";

import { getEvents } from "../../api/events";
import Person from "../../components/Person/Person";
import type { EventCardItemProps, PersonProps } from "../../components/Person/PersonProps";
import { useAppSelector } from "../../state/hooks";
import calculateAge from "../../utils/calculateAge";

type Event = Awaited<ReturnType<typeof getEvents>>[number];

export default function ViewProfile() {
  const userState = useAppSelector((state) => state.form);
  const events: EventCardItemProps[] = [];
  getEvents({ limit: 4, offset: 0 }).then((fetchedEvents: Event[]) => {
    fetchedEvents.forEach((fetchedEvent: Event) => {
      events.push({
        title: fetchedEvent.name,
        description: fetchedEvent.bio,
        imageUrl: fetchedEvent.images[0]
      });
    });
  });
  const user: PersonProps = {
    id: userState.id,
    firstName: userState.name,
    lastName: "No Last Name",
    age: calculateAge(new Date(userState.birthday)),
    pronouns: userState.pronouns,
    location: userState.location,
    school: userState.school,
    work: userState.job,
    height: {
      feet: userState.height.foot,
      inches: userState.height.inch
    },

    promptQuestion: userState.promptBio,
    promptResponse: userState.responseBio,
    interests: userState.passion,
    relationshipType: userState.looking,
    religion: userState.religion,
    politicalAffiliation: userState.politics,
    alchoholFrequency: userState.drinkHabit,
    smokingFrequency: userState.smokeHabit,
    drugFrequency: userState.drugHabit,
    cannabisFrequency: userState.weedHabit,
    instagramUsername: userState.instagram,
    mutualEvents: events,
    images: userState.photos.map((photo) => ({ image: photo.filePath, caption: photo.caption })),
    isMatched: true,
    likesYou: false,
    handleReact: () => {}
  };

  return (
    <View
      style={{
        flex: 1
      }}
    >
      <Text onPress={() => router.back()} style={scaledStyles.title}>{`< My Profile`}</Text>
      <Person
        id={user.id}
        firstName={user.firstName}
        lastName={user.lastName}
        age={user.age}
        pronouns={user.pronouns}
        location={user.location}
        school={user.school}
        work={user.work}
        height={user.height}
        promptQuestion={user.promptQuestion}
        promptResponse={user.promptResponse}
        interests={user.interests}
        relationshipType={user.relationshipType}
        religion={user.religion}
        politicalAffiliation={user.politicalAffiliation}
        alchoholFrequency={user.alchoholFrequency}
        smokingFrequency={user.smokingFrequency}
        drugFrequency={user.drugFrequency}
        cannabisFrequency={user.cannabisFrequency}
        instagramUsername={user.instagramUsername}
        mutualEvents={user.mutualEvents}
        images={user.images}
        isMatched={user.isMatched}
        likesYou={user.likesYou}
        handleReact={user.handleReact}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  title: {
    fontFamily: "DMSansBold",
    fontSize: 20,
    marginLeft: 20,
    marginTop: 40
  }
});

const scaledStyles = scaleStyleSheet(styles);
