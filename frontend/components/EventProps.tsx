export type EventCardProps = {
  id: number;
  title: string;
  description: string;
  date: string;
  price: number;
  location: string;
  handleReact: (like: boolean) => void;
};

export type EventImageCarouselProps = {
  images: string[];
};
