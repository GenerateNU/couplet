import * as ImagePicker from "expo-image-picker";
import * as MediaLibrary from "expo-media-library";
import React, { useState } from "react";
import { Image, TouchableOpacity, View, Text } from "react-native";
// import { RNS3 } from "react-native-aws3";
// import client from "../api/client";

interface PhotoPickerProps {
  onPick: (imgs: string[]) => void
}

export default function PhotoPicker({onPick}: PhotoPickerProps) {
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
      // client
      //   .PATCH("/users/{id}", {
      //     params: {
      //       path: { id: "5e91507e-5630-4efd-9fd4-799178870b11" }
      //     },
      //     body: {
      //       images: [
      //         "https://relay-file-upload.s3.amazonaws.com/06268d2f-715e-45b5-9a60-902e4bcc6456.jpg1710974218654"
      //       ],
      //       firstName: "karyna",
      //       lastName: "yen",
      //       age: 19
      //     }
      //   })
      //   .then((res) => {
      //     console.log("SUCCESS", res);
      //   })
      //   .catch((e) => {
      //     console.log("ERROR", e);
      //   });
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

    setImages(passedImages.map(img => img.uri))

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
    
    onPick(passedImages.map(img => img.uri))
  };

  const photoBoxStyling = {
    height: 150,
    width: 150,
    borderRadius: 10,
    borderWidth: 2,
    borderColor: "#CDCDCD",
    margin: 5
  };

  return (
    <View>
      <TouchableOpacity
        onPress={openPicker}
        style={{
          alignSelf: "center",
          width: 335,
          height: 335,
          justifyContent: "center",
          borderRadius: 10,
          borderStyle: "solid",
          borderWidth: 1,
          borderColor: "#CDCDCD",
        }}
      >
        <View
          style={{
            width: "100%",
            flexWrap: "wrap",
            flexDirection: "row",
            justifyContent: "center"
          }}
        >
          {[0, 1, 2, 3].map((i) =>
            i >= images.length ? (
              <View style={{ ...photoBoxStyling, borderStyle: "dashed", justifyContent: "center" }}>
                <Text style={{alignSelf: "center", fontSize: 50, color: "#CDCDCD"}}>+</Text>
              </View>
            ) : (
              <Image
                key={images[i]}
                source={{ uri: images[i] }}
                style={{ ...photoBoxStyling }}
              />
            )
          )}
        </View>
      </TouchableOpacity>
    </View>
  );
}
