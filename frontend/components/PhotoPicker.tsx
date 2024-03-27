import * as ImagePicker from "expo-image-picker";
import * as MediaLibrary from "expo-media-library";
import React, { useState } from "react";
import { Image, StyleSheet, TouchableOpacity, View } from "react-native";
import COLORS from "../colors";

const ADD_BUTTON = require("../assets/addbutton.png");

interface PhotoPickerProps {
  onPick: (imgs: string[]) => void;
}

export default function PhotoPicker({ onPick }: PhotoPickerProps) {
  const [images, setImages] = useState<string[]>([]);

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
    onPick(passedImages.map((img) => img.uri));
  };

  return (
    <View>
      <TouchableOpacity onPress={openPicker} style={styles.pressableContainer}>
        <View style={styles.pickerContainer}>
          {[0, 1, 2, 3].map((i) => (
            <View style={styles.photoContainer}>
              {i >= images.length ? (
                <View style={{ ...styles.photoBox, ...styles.emptyBox }}>
                  <Image source={ADD_BUTTON} style={styles.addButton} />
                </View>
              ) : (
                <Image key={images[i]} source={{ uri: images[i] }} style={styles.photoBox} />
              )}
            </View>
          ))}
        </View>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  pressableContainer: {
    alignSelf: "center",
    width: 335,
    height: 335,
    justifyContent: "center"
  },
  pickerContainer: {
    width: "100%",
    flexWrap: "wrap",
    flexDirection: "row",
    justifyContent: "center"
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
    borderRadius: 10,
    borderWidth: 2,
    borderColor: COLORS.darkGray,
    margin: 5
  },
  emptyBox: {
    borderStyle: "dashed",
    justifyContent: "center"
  },
  addButton: {
    position: "absolute",
    right: "-15%",
    bottom: "-15%"
  }
});
