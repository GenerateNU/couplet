import React from "react";
import { ScrollView, Text, View } from "react-native";
import MatchesUserCard from "./MatchesUserCard";

export type MatchesUser = {
    userID: number;
    name: string;
    birthday: number;
    location: string;
}

type MatchesUserSectionProps = {
    matches: MatchesUser[];
};

const MatchesUserSection: React.FC<MatchesUserSectionProps> = ({ matches }) => {
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