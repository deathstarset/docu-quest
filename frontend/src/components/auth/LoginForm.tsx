import { CardContent, CardFooter } from "@/components/ui/card";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import React, { useState } from "react";
import { z } from "zod";

const loginSchema = z.object({
  username: z.string().min(1, "Username is required"),
  password: z.string().min(8, "Password must be at least 6 characters long"),
});

export const LoginForm = () => {
  const [errors, setErrors] = useState<{
    username?: string;
    password?: string;
  }>({});

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    // getting the data
    const formData = new FormData(e.currentTarget);
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;

    // doing validation
    const validationResult = loginSchema.safeParse({
      username,
      password,
    });

    if (!validationResult.success) {
      const fieldErrors: {
        username?: string;
        password?: string;
      } = {};
      validationResult.error.errors.forEach((error) => {
        const field = error.path[0];
        const message = error.message;
        switch (field) {
          case "username":
            fieldErrors.username = message;
            break;
          case "password":
            fieldErrors.password = message;
            break;
        }
      });
      setErrors(fieldErrors);
      return;
    }

    setErrors({});
  }
  return (
    <form onSubmit={handleSubmit}>
      <CardContent className="flex flex-col gap-2">
        <div>
          <Label>Username</Label>
          <Input name="username" />
          {errors.username && (
            <span className="text-red-500">{errors.username}</span>
          )}
        </div>
        <div>
          <Label>Password</Label>
          <Input name="password" type="password" />
          {errors.password && (
            <span className="text-red-500">{errors.password}</span>
          )}
        </div>
      </CardContent>
      <CardFooter>
        <Button className="w-full">Login</Button>
      </CardFooter>
    </form>
  );
};
