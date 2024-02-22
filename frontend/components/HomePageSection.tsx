import { View, Text } from 'react-native'
import React from 'react'
import { ScrollView } from 'react-native'
import HomeEventCard from './HomeEventCard'
export default function HomePageSection({title, events} : {title: string, events: any[]}) {
  return (
    <View>
      <Text style={{fontSize: 20}}>
        {title} {" "}
      </Text>
      <View style={{flexDirection: "row"}}>
        <ScrollView horizontal={true}>
          {events.map((event, index) => {
            return <HomeEventCard key={index} />
          })}
        </ScrollView>
      </View>
    </View>
  )
}