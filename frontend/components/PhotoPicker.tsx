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

  useEffect(() => {
    onPick(images);
  }, [onPick, images]);

  const pick = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: false,
      allowsMultipleSelection: true,
      quality: 1,
      orderedSelection: true,
      selectionLimit: 4
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

    setImages(passedImages.map((img) => img.uri));
  };

  const removePhoto = (index: number) => {
    const newImages = [...images];
    newImages[index] = "";
    setImages(newImages);
  };

  return (
    <View>
      <View style={styles.pressableContainer}>
        {images.map((_, i) => (
          <TouchableOpacity onPress={openPicker} style={styles.photoContainer}>
            {images[i] === "" ? (
              <View style={{ ...styles.emptyBox }}>
                <Image source={ADD_BUTTON} style={styles.addRemoveButton} />
              </View>
            ) : (
              <View style={{ ...styles.photoBox }}>
                <Image
                  key={images[i]}
                  source={{ uri: images[i] }}
                  style={{ width: 140, height: 140, borderRadius: 10 }}
                />
                <Pressable onPress={() => removePhoto(i)} style={styles.addRemoveButton}>
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
    width: 335,
    height: 335,
    justifyContent: "center",
    flexWrap: "wrap",
    flexDirection: "row"
  },
  photoContainer: {
    height: 160,
    width: 160
  },
  photoBox: {
    height: 140,
    width: 140,
    left: 0,
    top: 0,
    borderRadius: 10
  },
  emptyBox: {
    height: 140,
    width: 140,
    left: 0,
    top: 0,
    borderRadius: 10,
    borderWidth: 2,
    borderColor: COLORS.darkGray,
    borderStyle: "dashed",
    justifyContent: "center"
  },
  addRemoveButton: {
    position: "absolute",
    right: "-15%",
    bottom: "-15%"
  }
});
