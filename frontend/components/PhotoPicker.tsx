import * as ImagePicker from "expo-image-picker";
import * as MediaLibrary from "expo-media-library";
import React, { useEffect, useState } from "react";
import { Image, Pressable, StyleSheet, TouchableOpacity, View } from "react-native";
import COLORS from "../colors";

const REMOVE_BUTTON = require("../assets/removebutton.png");
const ADD_BUTTON = require("../assets/addbutton.png");

interface PhotoPickerProps {
  onPick: (imgs: string[]) => void;
}

export default function PhotoPicker({ onPick }: PhotoPickerProps) {
  const [images, setImages] = useState<string[]>(["", "", "", ""]);
  const [imgCount, setImageCount] = useState<number>(0);

  useEffect(() => {
    onPick(images);

    setImageCount(images.filter((img) => img !== "").length);
  }, [onPick, images]);

  // The one issue with this is that it allows you to repeat the same photo
  // multiple times. Not sure how to persist the photos that you already have selected
  // in the picker, might need to look into it more.
  const pick = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: false,
      allowsMultipleSelection: true,
      quality: 1,
      orderedSelection: true,
      selectionLimit: 4 - imgCount
    });
    if (!result.canceled) {
      onDone(result.assets);
    }
  };

  const openPicker = async () => {
    const { status } = await MediaLibrary.getPermissionsAsync();
    if (status !== "granted") {
      // TODO: say we cant get their photos bc no permissions
      const newPerms = await MediaLibrary.requestPermissionsAsync();
      if (newPerms.granted) {
        pick();
      }
    } else {
      pick();
    }
  };

  const onDone = (passedImages: ImagePicker.ImagePickerAsset[]) => {
    if (typeof passedImages !== "object") return;
    if (!Object.prototype.hasOwnProperty.call(passedImages, "length")) return;

    const newImages = [...images];
    const toFill = passedImages.map((img) => img.uri);
    toFill.forEach((img) => {
      const nextToFill = newImages.findIndex((i) => i === "");
      newImages[nextToFill] = img;
    });

    setImages(newImages);
  };

  const removePhoto = (index: number) => {
    const newImages = [...images];
    newImages[index] = "";
    setImages(newImages);
  };

  return (
    <View>
      <View style={styles.pressableContainer}>
        {images.map((img, i) => (
          <TouchableOpacity onPress={openPicker} style={styles.photoContainer}>
            {img === "" ? (
              <View style={{ ...styles.emptyBox }}>
                <Image source={ADD_BUTTON} />
              </View>
            ) : (
              <View style={{ ...styles.photoBox }}>
                <Image
                  key={img}
                  source={{ uri: img }}
                  style={{ width: 140, height: 140, borderRadius: 10 }}
                />
                <Pressable onPress={() => removePhoto(i)} style={styles.removeButton}>
                  <Image source={REMOVE_BUTTON} />
                </Pressable>
              </View>
            )}
          </TouchableOpacity>
        ))}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  pressableContainer: {
    alignSelf: "center",
    width: 330,
    height: 330,
    justifyContent: "center",
    flexWrap: "wrap",
    flexDirection: "row"
  },
  photoContainer: {
    height: 140,
    width: 140,
    margin: 10
  },
  photoBox: {
    height: 140,
    width: 140,
    borderRadius: 10
  },
  emptyBox: {
    height: 140,
    width: 140,
    borderRadius: 10,
    backgroundColor: COLORS.disabled,
    justifyContent: "center",
    alignItems: "center"
  },
  removeButton: {
    position: "absolute",
    right: "-15%",
    bottom: "-15%"
  }
});
