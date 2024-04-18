import React from "react";
import { Dimensions, Image, View } from "react-native";
import Carousel from "react-native-reanimated-carousel";

export type EventImageCarouselProps = {
  images: string[];
};

function EventImageCarousel({ images }: EventImageCarouselProps) {
  const { height, width } = Dimensions.get("window");

  return (
    <View>
      <Carousel
        loop
        width={width}
        height={height / 2.5}
        autoPlay
        data={images}
        autoPlayInterval={3000}
        scrollAnimationDuration={1000}
        renderItem={({ index }) => (
          <View
            style={{
              justifyContent: "center"
            }}
          >
            <Image source={{ uri: images[index], width, height: height / 2.5 }} />
          </View>
        )}
      />
    </View>
  );
}

export default EventImageCarousel;
