import React, { useEffect, useState } from "react";
import { ActivityIndicator, Image, ScrollView, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import type { components } from "../../api/schema";
import getMatchesByUserId from "../../api/users";
import RECENT_NO_MATCHES from "../../assets/nomatches1.png";
import ALL_NO_MATCHES from "../../assets/nomatches2.png";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import LabelToggle from "../LabelToggle";

type User = components["schemas"]["User"];

export default function MatchesScreen() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [matches, setMatches] = useState<User[]>([]);
  const [matchFilter, setMatchFilter] = useState<string>("");
  const [displayMatches, setDisplayMatches] = useState<User[]>([]);
  const RECENT_NO_MATCHES_URI = Image.resolveAssetSource(RECENT_NO_MATCHES).uri;
  const ALL_NO_MATCHES_URI = Image.resolveAssetSource(ALL_NO_MATCHES).uri;

  useEffect(() => {
    const load = async () => {
      const res = await getMatchesByUserId("381edf55-cfd3-4651-985d-fdafffbb3002");
      setMatches([...res]);
    };

    setIsLoading(true);
    load();
  }, []);

  useEffect(() => {
    if (matchFilter === "All Matches") {
      setDisplayMatches(matches);
    } else if (matchFilter === "Recent") {
      // There is no Match struct, so there is no updatedAt or createdAt
    }
  }, [matchFilter, matches]);

  useEffect(() => {
    setIsLoading(false);
  }, [matches]);

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  return (
    <SafeAreaView style={scaledStyles.container}>
      <ScrollView
        contentContainerStyle={{ flexGrow: 1, paddingBottom: 100 }}
        showsVerticalScrollIndicator={false}
      >
        <View style={{ flexDirection: "column", alignItems: "center" }}>
          <View style={{ flexDirection: "column", alignSelf: "flex-start" }}>
            <Text style={scaledStyles.headingTitle}>Matches</Text>
            <LabelToggle labels={["Recent", "All Matches"]} onChange={setMatchFilter} />
            {displayMatches.length > 0 ? (
              <Text style={scaledStyles.headingText}>
                {matchFilter === "Recent"
                  ? "Your recent matches in the last week"
                  : "All of your past matches"}
              </Text>
            ) : null}
          </View>
        </View>

        {displayMatches.length > 0 ? (
          <View style={scaledStyles.matchesDisplay}>
            {displayMatches.map((match) => (
              <View style={scaledStyles.matchCard}>
                {match.images ? (
                  <Image style={scaledStyles.matchPhoto} source={{ uri: match.images[0] }} />
                ) : (
                  <View style={[scaledStyles.matchPhoto, { backgroundColor: COLORS.primary }]} />
                )}
                <Text style={scaledStyles.matchText}>
                  {match.firstName} {match.age}
                </Text>
              </View>
            ))}
          </View>
        ) : (
          <View style={scaledStyles.noMatchesDisplay}>
            {matchFilter === "Recent" ? (
              <>
                <Image
                  style={{ width: 300, height: 300 }}
                  source={{ uri: RECENT_NO_MATCHES_URI }}
                />
                <Text style={scaledStyles.noMatchesTitle}>No matches yet</Text>
                <Text style={scaledStyles.noMatchesText}>
                  Keep swiping to find your perfect plus one to your favorite events!
                </Text>
              </>
            ) : (
              <>
                <Image style={{ width: 300, height: 300 }} source={{ uri: ALL_NO_MATCHES_URI }} />
                <Text style={scaledStyles.noMatchesTitle}>No matches yet</Text>
                <Text style={scaledStyles.noMatchesText}>
                  Matches are made to have someone to go to events with. Go and like some people and
                  events!
                </Text>
              </>
            )}
          </View>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    height: "100%",
    margin: 24
  },
  headingTitle: {
    fontSize: 32,
    fontFamily: "DMSansBold",
    marginBottom: 12
  },
  headingText: {
    fontSize: 17,
    fontFamily: "DMSansMedium",
    marginTop: 16
  },
  matchesDisplay: {
    width: "100%",
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "space-between"
  },
  noMatchesDisplay: {
    marginTop: 16,
    width: "100%",
    justifyContent: "center",
    alignItems: "center"
  },
  noMatchesTitle: {
    fontSize: 28,
    fontFamily: "DMSansBold",
    textAlign: "center"
  },
  noMatchesText: {
    fontSize: 15,
    fontFamily: "DMSansMedium",
    textAlign: "center"
  },
  matchCard: {
    width: "47.5%",
    marginTop: 24,
    borderRadius: 8,
    backgroundColor: COLORS.white,
    shadowColor: COLORS.black,
    shadowRadius: 4,
    shadowOpacity: 0.1,
    shadowOffset: { width: 2, height: 4 }
  },
  matchPhoto: {
    height: 150,
    borderTopRightRadius: 8,
    borderTopLeftRadius: 8
  },
  matchText: {
    margin: 8,
    fontSize: 15,
    fontFamily: "DMSansMedium"
  }
});

const scaledStyles = scaleStyleSheet(styles);
