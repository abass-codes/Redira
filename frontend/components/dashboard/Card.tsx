interface Props {
    title: string;
    value: number;
  }
  
  export default function Card({
    title,
    value,
  }: Props) {
    return (
      <div className="rounded-xl border border-gray-800 bg-gray-900 p-6">
        <p className="text-sm text-gray-400">
          {title}
        </p>
  
        <h2 className="mt-2 text-4xl font-bold text-white">
          {value}
        </h2>
      </div>
    );
  }