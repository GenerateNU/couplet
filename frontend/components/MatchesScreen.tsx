import React, { useState } from "react";
import { View, Text, ScrollView } from "react-native";
import ConvosDropdown, { ConvoUser } from "./ConvosDropdown";
import MatchesUserSection, { MatchesUser } from "./MatchesUserSection";

// Dummy Data
const dummyActiveData: ConvoUser[] = [
  {
    userID: 1,
    name: "Barney",
    chatLine: "Mark your calendars!"
  },
  {
    userID: 2,
    name: "Camey",
    chatLine: "Here is my number"
  },
  {
    userID: 3,
    name: "Camey",
    chatLine: "How is this event?"
  },
];

const dummyArchiveData: ConvoUser[] = [
  {
    userID: 1,
    name: "Jay",
    chatLine: "Sounds good"
  },
  {
    userID: 2,
    name: "Liam",
    chatLine: "See you soon"
  },
  {
    userID: 3,
    name: "Perval",
    chatLine: "Bye"
  },
  {
    userID: 4,
    name: "Todd",
    chatLine: "Okay, see you soon"
  },
  {
    userID: 5,
    name: "Ellie",
    chatLine: "See you Later"
  },
  {
    userID: 6,
    name: "Adam",
    chatLine: "Bye"
  },
  {
    userID: 7,
    name: "Reyna",
    chatLine: "Sounds good"
  },
];

const dummyUserData: MatchesUser[] = [
  {
    userID: 1,
    name: "Arnold",
    birthday: 21,
    location: "Boston",
  },
  {
    userID: 2,
    name: "Bob",
    birthday: 25,
    location: "New York",
  },
  {
    userID: 3,
    name: "Dan",
    birthday: 29,
    location: "Boston",
  },
  {
    userID: 4,
    name: "John",
    birthday: 18,
    location: "New York",
  },
  {
    userID: 5,
    name: "Tom",
    birthday: 22,
    location: "Boston",
  },
]

export default function MatchesScreen() {
    
  // Code for figuring out if they have any Likes
  // Using Dummy Data currently
  const [activeConvos, setActiveConvos] = useState(dummyActiveData);
  const [archiveConvos, setArchiveConvos] = useState(dummyArchiveData);
  const [activeUsers, setActiveUsers] = useState(dummyUserData);

  return (
    <ScrollView contentContainerStyle={{ flexGrow: 1, paddingBottom: 100 }} showsVerticalScrollIndicator={false}>
      <View style={{ flexDirection: 'column', alignItems: 'center' }}>
        <View style={{ borderBottomWidth: 1, borderBottomColor: 'black', width: '100%', alignItems: 'center' }}>
          <Text style={{ marginTop: 56, fontSize: 32, marginBottom: 9 }}>Matches</Text>
        </View>
        <View style={{ flex: 1, width: '100%'}}>
          <Text style={{ fontSize: 22, lineHeight: 32, paddingLeft: '5%' }}>
            Make the first move!
          </Text>
          <MatchesUserSection matches={activeUsers} />
          <ConvosDropdown convos={activeConvos} convoType="activeConvos" />
          <ConvosDropdown convos={archiveConvos} convoType="archiveConvos" />
        </View>
      </View>
    </ScrollView>
  );
}