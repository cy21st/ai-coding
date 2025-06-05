package com.example.controller;

import com.example.common.ApiResponse;
import com.example.dto.LoginRequest;
import com.example.service.AdminUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api")
public class AuthController {

    @Autowired
    private AdminUserService adminUserService;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @PostMapping("/login")
    public ApiResponse<?> login(@RequestBody LoginRequest loginRequest) {
        return adminUserService.login(loginRequest);
    }

    @PostMapping("/logout")
    public ApiResponse<Map<String, String>> logout() {
        return ApiResponse.success("Successfully logged out", 
            Map.of("message", "Successfully logged out"));
    }

    @GetMapping("/test")
    public ApiResponse<String> test() {
        String hashedPassword = passwordEncoder.encode("123456");
        System.out.println("Java bcrypt hash for '123456': " + hashedPassword);
        return ApiResponse.success(hashedPassword);
    }
} 