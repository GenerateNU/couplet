import React from "react";
import { ScrollView, StyleSheet, Text, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";
import HomeEventCard from "./HomeEventCard";

const MemoizedHomeEventCard = React.memo(HomeEventCard);

export default function HomePageSection({ title, events }: { title: string; events: any[] }) {
  return (
    <View style={{ marginTop: 32 }}>
      <Text style={scaledStyles.sectionTitle}>{`${title} >`} </Text>
      <View style={{ flexDirection: "row" }}>
        <ScrollView horizontal>
          {events.map((event) => (
            <MemoizedHomeEventCard key={event.id} id={event.id} name={event.name} />
          ))}
        </ScrollView>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  sectionTitle: {
    fontSize: 20,
    fontWeight: "700",
    fontFamily: "DMSansMedium",
    letterSpacing: -0.2,
    marginBottom: 8
  }
});

const scaledStyles = scaleStyleSheet(styles);
