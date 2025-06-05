package com.example.common;

import lombok.Data;

@Data
public class ApiResponse<T> {
    private int code;
    private String message;
    private T data;
    private String error;

    public static <T> ApiResponse<T> success(T data) {
        ApiResponse<T> response = new ApiResponse<>();
        response.setCode(200);
        response.setMessage("Success");
        response.setData(data);
        return response;
    }

    public static <T> ApiResponse<T> success(String message, T data) {
        ApiResponse<T> response = new ApiResponse<>();
        response.setCode(200);
        response.setMessage(message);
        response.setData(data);
        return response;
    }

    public static <T> ApiResponse<T> fail(int code, String message, String error) {
        ApiResponse<T> response = new ApiResponse<>();
        response.setCode(code);
        response.setMessage(message);
        response.setError(error);
        return response;
    }

    public static <T> ApiResponse<T> badRequest(String message, String error) {
        return fail(400, message, error);
    }

    public static <T> ApiResponse<T> unauthorized(String message, String error) {
        return fail(401, message, error);
    }

    public static <T> ApiResponse<T> notFound(String message, String error) {
        return fail(404, message, error);
    }

    public static <T> ApiResponse<T> internalError(String message, String error) {
        return fail(500, message, error);
    }
} 