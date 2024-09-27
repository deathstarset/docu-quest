import React from "react";

interface TypographyH2Props extends React.ComponentProps<"h2"> {
  children: React.ReactNode;
  className?: string;
}

export const TypographyH2 = ({
  children,
  className,
  ...props
}: TypographyH2Props) => {
  return (
    <h2
      className={`scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0 ${className}`}
      {...props}
    >
      {children}
    </h2>
  );
};
