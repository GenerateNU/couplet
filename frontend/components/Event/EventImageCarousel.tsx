import React from "react";
import { Dimensions, Image, View } from "react-native";
import Carousel from "react-native-reanimated-carousel";

export type EventImageCarouselProps = {
  images: string[];
};

function EventImageCarousel({ images }: EventImageCarouselProps) {
  const { width } = Dimensions.get("window");

  return (
    <View>
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
            <Image source={{ uri: images[index], height: width }} />
          </View>
        )}
      />
    </View>
  );
}

export default EventImageCarousel;
