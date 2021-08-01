import { Component, ErrorInfo, ReactNode } from "react";

export interface ErrorBoundaryProps {
  children: ReactNode;
  onError?: (error: Error, errorInfo: ErrorInfo) => void;
}

interface ErrorBoundaryState {
  error: Error | null;
}

export class ErrorBoundary extends Component<
  ErrorBoundaryProps,
  ErrorBoundaryState
> {
  public static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    return { error };
  }

  public state: ErrorBoundaryState = {
    error: null,
  };

  public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    if (this.props.onError) {
      this.props.onError(error, errorInfo);
    }
  }

  public render() {
    if (this.state.error !== null) {
      return <h1>{this.state.error.message}</h1>;
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
