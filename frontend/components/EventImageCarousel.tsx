import React from "react";
import { Image, View, Dimensions } from "react-native";
import Carousel from 'react-native-reanimated-carousel';

function EventImageCarousel() {
  const dummyImages = [
    "https://marvel-b1-cdn.bc0a.com/f00000000283318/home.dartmouth.edu/sites/home/files/styles/max_width_720px/public/2023-12/20220127_around_campus_eb_157.jpg?itok=bJJ9L7nZ",
    "https://www.lawnstarter.com/blog/wp-content/uploads/2022/12/iStock-1423384637-2-feature-image-1.jpg", 
    "https://www.flightonice.com/wp-content/uploads/2022/10/e4d4996c-da07-403e-a1c9-17696615d7ea_750x422.jpg", 
    "https://www.novaparks.com/sites/default/files/styles/scale_1440/public/2024-01/IceSkating202312190151_NP.jpg?itok=a6ScPTLd"
  ];
  const {width} = Dimensions.get('window');
  return (
    <View
      style={{
        width: "100%",
        borderRadius: 12,
        alignSelf: "center",
        height: "32%"
      }}
    >
            <Carousel
                loop
                width={width}
                height={width}
                autoPlay
                data={ dummyImages }
                scrollAnimationDuration={1000}
                renderItem={({ index }) => (
                    <View
                        style={{
                            flex: 1,
                            justifyContent: 'center',
                        }}
                    >
                        <Image source={{uri: dummyImages[index], height: 300}} /> 
                    </View>
                )}
            />
    </View>
  );
}

export default EventImageCarousel;
