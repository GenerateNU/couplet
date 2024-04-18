import React from "react";
import { View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import PeopleStack from "../components/Person/PeopleStack";

export default function People() {
  return (
    // TODO: GET THE ID
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <PeopleStack userId="1234" />
      <Navbar activePage="" />
    </View>
  );
}
