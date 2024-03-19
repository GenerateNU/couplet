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
      <EventCard
        id={1}
        title="Winter Ice Skating"
        description="The best place to skate outdoors in Boston. Whether you are a first-time skater or an aspiring Olympian, Frog Pond enables you to take pleasure in the sport of ice skating. Frog Pond offers public ice skating, skate rentals, and skate sharpening."
        date="1/2/24"
        price={20}
        location="Frog Pond"
        handleReact={handleReact}
      />,
      <EventCard
        id={2}
        title="Movie"
        description="Watch a movie at the theater"
        date="2/2/24"
        price={40}
        location="AMC"
        handleReact={handleReact}
      />,
      <EventCard
        id={3}
        title="Hockey"
        description=""
        date="3/2/24"
        price={30}
        location="TD Garden"
        handleReact={handleReact}
      />
    ];
    cardLength.current = dummyStack.length;
    setCards(dummyStack);
  }, [handleReact]);

  const CurrentCard = cards[currentCardIndex];

  return <View>{CurrentCard}</View>;
}

export default CardStack;
