package com.example.service;

import com.example.common.ApiResponse;
import com.example.dto.LoginRequest;
import com.example.entity.AdminUser;
import com.example.repository.AdminUserRepository;
import com.example.security.JwtTokenProvider;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

@Service
public class AdminUserService implements UserDetailsService {

    private final AdminUserRepository adminUserRepository;
    private final ApplicationContext applicationContext;
    private final JwtTokenProvider jwtTokenProvider;

    @Autowired
    public AdminUserService(AdminUserRepository adminUserRepository,
                          ApplicationContext applicationContext,
                          JwtTokenProvider jwtTokenProvider) {
        this.adminUserRepository = adminUserRepository;
        this.applicationContext = applicationContext;
        this.jwtTokenProvider = jwtTokenProvider;
    }

    private AuthenticationManager getAuthenticationManager() {
        return applicationContext.getBean(AuthenticationManager.class);
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        AdminUser adminUser = adminUserRepository.findByUsername(username)
                .orElseThrow(() -> new UsernameNotFoundException("User not found: " + username));

        return new User(adminUser.getUsername(),
                       adminUser.getPassword(),
                       Collections.singletonList(new SimpleGrantedAuthority("ROLE_" + adminUser.getRole().toUpperCase())));
    }

    public ApiResponse<Map<String, Object>> login(LoginRequest request) {
        try {
            Authentication authentication = getAuthenticationManager().authenticate(
                new UsernamePasswordAuthenticationToken(request.getUsername(), request.getPassword())
            );

            AdminUser user = adminUserRepository.findByUsername(request.getUsername())
                .orElseThrow(() -> new RuntimeException("User not found"));

            String token = jwtTokenProvider.generateToken(authentication);

            Map<String, Object> userInfo = new HashMap<>();
            userInfo.put("id", user.getId());
            userInfo.put("username", user.getUsername());
            userInfo.put("role", user.getRole());

            Map<String, Object> response = new HashMap<>();
            response.put("token", token);
            response.put("user", userInfo);

            return ApiResponse.success("Login successful", response);
        } catch (Exception e) {
            return ApiResponse.unauthorized("Invalid username or password", e.getMessage());
        }
    }
} 