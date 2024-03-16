import React, { useCallback, useEffect, useRef, useState } from "react";
import { View } from "react-native";
import EventCard from "./EventCard";

function CardStack() {
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [cards, setCards] = useState<React.JSX.Element[]>([]);
  const cardLength = useRef(cards.length);

  const handleReact = useCallback((like: boolean): boolean => {
    setCurrentCardIndex((prevIndex) => (prevIndex + 1) % cardLength.current);
    return like;
    // TODO: Add logic to update the database with the user's reaction to the event
  }, []);

  useEffect(() => {
    const dummyStack = [
      <EventCard id="123e4567-e89b-12d3-a456-426614174000" handleReact={handleReact} />,
      <EventCard id="123e4567-e89b-12d3-a456-426614174000" handleReact={handleReact} />,
      <EventCard id="123e4567-e89b-12d3-a456-426614174000" handleReact={handleReact} />
    ];
    cardLength.current = dummyStack.length;
    setCards(dummyStack);
  }, [handleReact]);

  const CurrentCard = cards[currentCardIndex];

  return <View>{CurrentCard}</View>;
}

export default CardStack;
