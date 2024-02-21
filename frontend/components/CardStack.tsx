import React, { useState } from "react";
import { View } from "react-native";
import EventCard from "./EventCard";

const CardStack = () => {
  const [currentCardIndex, setCurrentCardIndex] = useState(0);

  function handleReact(like: Boolean): void {
    console.log("likes:" + like);
    setCurrentCardIndex((prevIndex) => (prevIndex + 1) % dummyStack.length);
    console.log("new Card Switched!");

    // TODO: Add logic to update the database with the user's reaction to the event
  }

  const dummyStack = [
    <EventCard
      id={1}
      title={"Museum"}
      description="Night at the museum"
      date="1/2/24"
      location="MFA"
      handleReact={handleReact}
    />,
    <EventCard
      id={2}
      title={"Movie"}
      description="Watch a movie at the theater"
      date="2/2/24"
      location="AMC"
      handleReact={handleReact}
    />,
    <EventCard
      id={3}
      title={"Hockey"}
      description="Bruins game"
      date="3/2/24"
      location="TD Garden"
      handleReact={handleReact}
    />
  ];
  const CurrentCard = dummyStack[currentCardIndex];

  return <View>{CurrentCard}</View>;
};

export default CardStack;
