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
    setImages([]);

    if (typeof passedImages !== "object") return;
    if (!Object.prototype.hasOwnProperty.call(passedImages, "length")) return;

    passedImages.forEach((img) => {
      console.log("hi", img.fileName);
      setImages((imgs) => [...imgs, img.uri]);
    });

    // passedImages.forEach(async (img) => {
    //   let assetInfo;
    //   if (Object.prototype.hasOwnProperty.call(img, "assetId") && img.assetId != null)
    //     assetInfo = await MediaLibrary.getAssetInfoAsync(img.assetId);
    //   else return;
    //   if (assetInfo.localUri == null || img.fileName == null) return;
    //   const extension = assetInfo.localUri.substring(assetInfo.localUri.lastIndexOf(".") + 1);

    //   const type = `${img.type}/${extension.toLowerCase()}`;
    //   const uri = assetInfo.localUri;
    //   const name = img.fileName + new Date().getTime();

    //   const file = {
    //     uri,
    //     name,
    //     type
    //   };

    //   const options = {
    //     bucket: "relay-file-upload",
    //     region: "us-east-2",
    //     accessKey: process.env.EXPO_PUBLIC_AWS_ACCESS_KEY_ID || "",
    //     secretKey: process.env.EXPO_PUBLIC_AWS_SECRET_ACCESS_KEY || "",
    //     successActionStatus: 201
    //   };

    //   RNS3.put(file, options)
    //     .then((res) => {
    //       if (res.status !== 201) throw new Error("Failed to upload image to S3");
    //       // We uploaded it yay! Now we can do something with the URL
    //       // @ts-ignore
    //       console.log(res.body.postResponse.location);
    //       // @ts-ignore
    //       setImages([...images, res.body.postResponse.location]);
    //       // TODO: Backend call with the image we just uploaded
    //     })
    //     .catch((e) => {
    //       console.log(e);
    //     });
    // });

    // fetch(`http://${process.env.BACKEND_ADDRESS}/users/050565f3-f71d-4baa-9dcc-d6d822f03dd6`, {
    //   method: "PATCH",
    //   body: JSON.stringify({ images })
    // }).catch((e) => {
    //   console.log(e);
    // });
  };
  return (
    <ScrollView>
      <Text>PhotoPicker</Text>
      <TouchableOpacity
        onPress={openPicker}
        style={{
          width: 300,
          height: 100,
          backgroundColor: "white",
          justifyContent: "center",
          borderRadius: 10,
          borderStyle: "solid",
          borderWidth: 1
        }}
      >
        <Text
          style={{
            color: "black",
            fontSize: 24,
            textAlign: "center"
          }}
        >
          upload !!!!!!!
        </Text>
      </TouchableOpacity>

      <Text>Image:</Text>
      <View>
        {images.map((img) => (
          <Image key={img} source={{ uri: img }} style={{ width: 300, height: 300 }} />
        ))}
      </View>
    </ScrollView>
  );
}

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
