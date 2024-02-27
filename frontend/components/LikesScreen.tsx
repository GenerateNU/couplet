import React /*{useState}*/ from "react";
import { View, Text, Image } from "react-native";

const LIKE_IMAGE = require("../assets/likes.png");

export default function LikesScreen() {

  // Code for figuring out if they have any Likes
  // const [likes, setLikes] = useState([]);

  return (
    <View style={{ flexDirection: 'column', alignItems: 'center' }}>
      <View style={{ borderBottomWidth: 1, borderBottomColor: 'black', width: '100%', alignItems: 'center' }}>
        <Text style={{ marginTop: 56, fontSize: 32, marginBottom: 9 }}>Likes</Text>
      </View>
      <View style={{ alignItems: 'center', height: '100%', width: '100%'}}>
        <Image source={LIKE_IMAGE} style={{ width: 131, height: 131, marginTop: 100, marginBottom: 31}} />
        <Text style={{ width: '75%', fontSize: 14, lineHeight: 24, textAlign: 'center'}}> 
          Seems like you don&apos;t have any likes yet. Go out there and start liking some people. 
        </Text>
      </View>
    </View>
  );
}
