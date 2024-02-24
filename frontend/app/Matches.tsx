import React from "react";
import { createStackNavigator } from '@react-navigation/stack';

const MatchStack = createStackNavigator();

export default function Matches() {
  return (
    <MatchStack.Navigator>
      <MatchStack.Screen name="Home" component={HomeScreen} />
      <MatchStack.Screen name="Notifications" component={NotificationsScreen} />
      <MatchStack.Screen name="Profile" component={ProfileScreen} />
      <MatchStack.Screen name="Settings" component={SettingsScreen} />
    </MatchStack.Navigator>
  );
}
