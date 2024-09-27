import React from "react";

interface TypographyH1Props extends React.ComponentProps<"h1"> {
  children: React.ReactNode;
  className?: string;
}

export const TypographyH1 = ({
  children,
  className,
  ...props
}: TypographyH1Props) => {
  return (
    <h1
      className={`scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl ${className}`}
      {...props}
    >
      {children}
    </h1>
  );
};
