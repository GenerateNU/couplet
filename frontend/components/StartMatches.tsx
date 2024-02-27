import React from "react";
import { Dimensions, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import Carousel, { ICarouselInstance } from "react-native-reanimated-carousel";
import { SafeAreaView } from "react-native-safe-area-context";
import { EventCardItem } from "./EventCardItem";

export default function StartMatches() {
  const [data, setData] = React.useState([...new Array(6).keys()]); // Change this data to match users events
  const ref = React.useRef<ICarouselInstance>(null);
  const PAGE_WIDTH = Dimensions.get("window").width;

  const today = new Date();
  const date = `${today.toLocaleString("default", {
    month: "long"
  })} ${today.getDate()}th, ${today.getFullYear()}`;

  setData(data.map((_, index) => index));

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <View style={styles.container}>
        <Text style={styles.dateText}>{date}</Text>
        <View style={styles.notificationContainer}>
          <Text style={styles.notificationText}>You matched!{"\n"}Here is Arnold’s #</Text>
        </View>
        <View style={styles.phoneNumberContainer}>
          <Text style={styles.phoneNumber}>(617)-111-1111</Text>
        </View>
        <Text style={styles.suggestDateText}>Suggest a first date!</Text>
      </View>
      <Carousel
        vertical={false}
        width={PAGE_WIDTH / 2.5}
        height={PAGE_WIDTH / 1.5}
        ref={ref}
        style={{ width: "100%" }}
        data={data}
        pagingEnabled
        renderItem={({ index }) => (
          <EventCardItem
            title={`Title ${index + 1}`}
            description={`Description ${index + 1}`}
            imageUrl={`https://example.com/image${index + 1}.jpg`}
          />
        )}
      />
      <TouchableOpacity style={styles.browseMoreButton}>
        <Text style={styles.browseMoreButtonText}>Browse more</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#FFF",
    alignItems: "center"
  },
  dateText: {
    color: "#898A8D",
    marginTop: 29
  },
  matchMessage: {
    textAlign: "center",
    marginTop: 30,
    fontFamily: "Comic Neue",
    fontSize: 24
  },
  phoneNumberContainer: {
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 10,
    borderWidth: 1,
    marginTop: 15,
    padding: 8
  },
  phoneNumber: {
    textAlign: "center",
    fontFamily: "Comic Neue, sans-serif",
    fontSize: 18,
    padding: 8
  },
  suggestDateText: {
    marginTop: 100,
    paddingLeft: 22,
    fontFamily: "Comic Neue, sans-serif",
    fontSize: 22,
    paddingBottom: 10
  },
  browseMoreButton: {
    backgroundColor: "#000",
    borderRadius: 10,
    padding: 16,
    marginLeft: 22,
    maxWidth: 150
  },
  browseMoreButtonText: {
    color: "#FFF",
    textAlign: "center",
    fontFamily: "Comic Neue, sans-serif",
    fontSize: 18
  },
  notificationContainer: {
    backgroundColor: "#fff",
    padding: 20,
    borderRadius: 8,
    alignItems: "center",
    justifyContent: "center"
  },
  notificationText: {
    color: "#000",
    textAlign: "center",
    letterSpacing: 1.2,
    fontSize: 24,
    lineHeight: 44,
    fontFamily: "Comic Neue, sans-serif"
  }
});
