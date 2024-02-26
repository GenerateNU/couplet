import React from "react";
import { Dimensions, Image, View } from "react-native";
import Carousel from "react-native-reanimated-carousel";
import { EventImageCarouselProps } from "./EventProps";



function EventImageCarousel({ images }: EventImageCarouselProps) {
  const { width } = Dimensions.get("window");
  return (
    <View
      style={{
        width: "100%",
        borderRadius: 12,
        alignSelf: "center",
        height: "31%"
      }}
    >
      <Carousel
        loop
        width={width}
        height={width}
        autoPlay
        data={images}
        scrollAnimationDuration={1000}
        renderItem={({ index }) => (
          <View
            style={{
              justifyContent: "center"
            }}
          >
            <Image source={{ uri: images[index], height: 350 }} />
          </View>
        )}
      />
    </View>
  );
}

export default EventImageCarousel;
