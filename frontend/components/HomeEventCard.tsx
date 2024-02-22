import { View, Text, Image } from 'react-native'
import React from 'react'
export default function HomeEventCard() {
  return (
    <View style={{
        borderStyle: 'solid',
        borderWidth: 1,
        borderColor: 'black',
        borderRadius: 10,
        borderTopLeftRadius: 0,
        borderTopRightRadius: 0,
        marginRight: 10,

    }}>
      {/* <Image source={{uri: 'https://reactnative.dev/img/tiny_logo.png'}} style={{width: 150, height: 150}} /> */}
        <View style={{width:"100%", height: 150,backgroundColor:"rgb(200,200,200)"}}></View>
        <View>
            <Text style={{textAlign: "center", padding: 10, fontSize: 14}}>Winter Ice Skating</Text>
            <View style={{flexDirection: "row", padding: 10, borderRadius: 20, paddingTop:0}}>
                <Image source={require("../assets/pin.png")} style={{width: 20, height: 20}} /> 
                <Text style={{textAlign: "center", justifyContent: "center", verticalAlign: "middle", marginTop: 2}}>Frog Pond</Text>
                <Image source={require("../assets/coin.png")} style={{width: 20, height: 20, marginLeft: 10}} /> 
                <Text style={{textAlign: "center", justifyContent: "center", verticalAlign: "middle", marginTop: 2, marginHorizontal: 2}}>Cost</Text>
            <View />
        </View>
        </View>
    </View>
  )
}