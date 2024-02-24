import React from "react";
import { Image, Text, View } from "react-native";

function EventImageCarousel() {
  const dummyImages = [
    "https://cdn.britannica.com/55/174255-050-526314B6/brown-Guernsey-cow.jpg",
    "https://cff2.earth.com/uploads/2023/05/16064103/Farms-960x640.jpg"
  ];
  return (
    <View
      style={{
        width: "100%",
        borderRadius: 12,
        alignSelf: "center",
        marginBottom: 10,
        height: "40%"
      }}
    >
      <Text>Event Image Carousel</Text>
      <Image source={{ uri: dummyImages[0], height: 300 }} />
      {/* <Image
                source={{uri: dummyImages[0]}}
            /> */}
    </View>
  );
}

export default EventImageCarousel;
