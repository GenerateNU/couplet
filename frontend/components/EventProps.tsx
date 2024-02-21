export interface EventCardProps {
  id: number;
  title: string;
  description: string;
  date: string;
  location: string;

  handleReact: (like: boolean) => void;
}
