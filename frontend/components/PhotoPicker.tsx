import * as ImagePicker from "expo-image-picker";
import * as MediaLibrary from "expo-media-library";
import React, { useState } from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { RNS3 } from "react-native-aws3";

export default function PhotoPicker() {
  const [images, setImages] = useState<string[]>([]);

  const pick = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: false,
      allowsMultipleSelection: true,
      quality: 1,
      orderedSelection: true
    });
    console.log(result);
    if (!result.canceled) {
      onDone(result.assets);
    }
  };
  const openPicker = async () => {
    const { status } = await MediaLibrary.getPermissionsAsync();
    console.log(status);
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

    passedImages.forEach(async (img) => {
      let assetInfo;
      if (Object.prototype.hasOwnProperty.call(img, "assetId") && img.assetId != null)
        assetInfo = await MediaLibrary.getAssetInfoAsync(img.assetId);
      else return;
      if (assetInfo.localUri == null || img.fileName == null) return;
      const extension = assetInfo.localUri.substring(assetInfo.localUri.lastIndexOf(".") + 1);

      const type = `${img.type}/${extension.toLowerCase()}`;
      const uri = assetInfo.localUri;
      const name = img.fileName + new Date().getTime();

      const file = {
        uri,
        name,
        type
      };

      const options = {
        bucket: "relay-file-upload",
        region: "us-east-2",
        accessKey: process.env.EXPO_PUBLIC_AWS_ACCESS_KEY_ID || "",
        secretKey: process.env.EXPO_PUBLIC_AWS_SECRET_ACCESS_KEY || "",
        successActionStatus: 201
      };

      RNS3.put(file, options)
        .then((res) => {
          if (res.status !== 201) throw new Error("Failed to upload image to S3");
          // We uploaded it yay! Now we can do something with the URL
          // @ts-ignore
          console.log(res.body.postResponse.location);
          // @ts-ignore
          setImages([...images, res.body.postResponse.location]);
          // TODO: Backend call with the image we just uploaded
        })
        .catch((e) => {
          console.log(e);
        });
    });

    fetch(`http://${process.env.BACKEND_ADDRESS}/users/050565f3-f71d-4baa-9dcc-d6d822f03dd6`, {
      method: "PATCH",
      body: JSON.stringify({ images })
    }).catch((e) => {
      console.log(e);
    });
  };
  return (
    <View>
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
      {/* {images && images.map((i) => (
        <Image src={i} style={{ width: 300, height: 300 }} />
      ))} */}
    </View>
  );
}
