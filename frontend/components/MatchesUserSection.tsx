import React from "react";
import { ScrollView, View } from "react-native";
import MatchesUserCard from "./MatchesUserCard";

export type MatchesUser = {
  userID: number;
  name: string;
  birthday: number;
  location: string;
};

type MatchesUserSectionProps = {
  matches: MatchesUser[];
};

function MatchesUserSection({ matches }: MatchesUserSectionProps) {
  return (
    <View style={{ marginVertical: 10, marginLeft: 10 }}>
      <View style={{ flexDirection: "row" }}>
        <ScrollView horizontal showsHorizontalScrollIndicator={false}>
          {matches.map((user) => (
            <MatchesUserCard key={user.userID} profile={user} />
          ))}
        </ScrollView>
      </View>
    </View>
  );
}

export default MatchesUserSection;
