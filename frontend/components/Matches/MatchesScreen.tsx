import React, { useEffect, useState } from "react";
import { ActivityIndicator, Image, ScrollView, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import type { paths } from "../../api/schema";
import getMatchesByUserId from "../../api/users";
import RECENT_NO_MATCHES from "../../assets/nomatches1.png";
import ALL_NO_MATCHES from "../../assets/nomatches2.png";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import LabelToggle from "../LabelToggle";

type Matches = paths["/matches/{id}"]["get"]["responses"][200]["content"]["application/json"];

const SEVEN_DAYS = 604800000;

export default function MatchesScreen() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [matches, setMatches] = useState<Matches>([]);
  const [matchFilter, setMatchFilter] = useState<string>("Recent");
  const [displayMatches, setDisplayMatches] = useState<Matches>([]);
  const RECENT_NO_MATCHES_URI = Image.resolveAssetSource(RECENT_NO_MATCHES).uri;
  const ALL_NO_MATCHES_URI = Image.resolveAssetSource(ALL_NO_MATCHES).uri;

  useEffect(() => {
    const load = async () => {
      // TODO: Change to own user id
      const userId = "0f46ee58-5577-4b83-99b7-fff1b8de0e1f";
      const res = await getMatchesByUserId(userId);
      setMatches([...res]);
    };

    setIsLoading(true);
    load();
  }, []);

  useEffect(() => {
    if (matchFilter === "All Matches") {
      setDisplayMatches(matches);
    } else if (matchFilter === "Recent") {
      setDisplayMatches(
        matches.filter((m) => Date.now().valueOf() - new Date(m.createdAt).valueOf() < SEVEN_DAYS)
      );
    }
  }, [matchFilter, matches]);

  useEffect(() => {
    setIsLoading(false);
  }, [matches]);

  if (isLoading) {
    return (
      <View style={{ flex: 1, height: "100%", justifyContent: "center", alignItems: "center" }}>
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
        <View style={scaledStyles.headingContainer}>
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

        {displayMatches.length > 0 ? (
          <View style={scaledStyles.matchesDisplay}>
            {displayMatches.map((match) => (
              <View key={match.id} style={scaledStyles.matchCard}>
                {match.images ? (
                  <Image style={scaledStyles.matchPhoto} source={{ uri: match.images[0] }} />
                ) : (
                  <View style={[scaledStyles.matchPhoto, { backgroundColor: COLORS.primary }]} />
                )}
                <View style={scaledStyles.matchTextContainer}>
                  <Text style={scaledStyles.matchNotViewedDot}>
                    {!match.viewed ? "\u2B24" : null}{" "}
                  </Text>
                  <Text style={[scaledStyles.matchText, { fontFamily: "DMSansBold" }]}>
                    {match.firstName}
                  </Text>
                  <Text style={[scaledStyles.matchText, { fontFamily: "DMSansMedium" }]}>
                    {match.age}
                  </Text>
                </View>
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
  headingContainer: {
    flexDirection: "column",
    alignSelf: "flex-start"
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
    marginVertical: 8,
    marginHorizontal: 2,
    fontSize: 15
  },
  matchNotViewedDot: {
    color: COLORS.primary,
    fontSize: 8
  },
  matchTextContainer: {
    marginHorizontal: 8,
    flexDirection: "row",
    alignItems: "center"
  }
});

const scaledStyles = scaleStyleSheet(styles);
