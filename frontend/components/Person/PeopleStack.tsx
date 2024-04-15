import React, { useCallback, useEffect, useState } from "react";
import { View } from "react-native";
import { ActivityIndicator } from "react-native-paper";
import { getUsers } from "../../api/users";
import Navbar from "../Layout/Navbar";
import Person from "./Person";

export default function PeopleStack() {
  const [people, setPeople] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleReact = useCallback(
    (like: boolean) => {
      // TODO - find event swipe function
      // personSwipe(userId, currentPersonId, like).then()

      // pop the current user from the queue
      setPeople([...people].filter((_, i) => i !== 0));
    },
    [people]
  );

  // Load users to put in card stack
  useEffect(() => {
    const load = async () => {
      setIsLoading(true);
      try {
        // Will have to change this function to the user recommendation endpoint
        // Needs a boolean in th response to indicate whether they liked you
        const res = await getUsers({ limit: 10 });
        setPeople(res.map((p) => p.id));
      } catch (error) {
        throw new Error(String(error));
      }
      setIsLoading(false);
    };

    load();
  }, []);

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Person
        id={people[0]}
        isMatched={false}
        likesYou // placeholder
        handleReact={handleReact}
      />

      <Navbar activePage="" />
    </View>
  );
}
