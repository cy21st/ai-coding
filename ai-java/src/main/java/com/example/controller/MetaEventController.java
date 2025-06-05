package com.example.controller;

import com.example.common.ApiResponse;
import com.example.entity.MetaEvent;
import com.example.service.MetaEventService;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/api/events")
public class MetaEventController {

    private final MetaEventService metaEventService;

    public MetaEventController(MetaEventService metaEventService) {
        this.metaEventService = metaEventService;
    }

    @GetMapping
    public ApiResponse<Map<String, Object>> findAll(
            @RequestParam(defaultValue = "1") int pageNum,
            @RequestParam(defaultValue = "10") int pageSize) {
        Page<MetaEvent> events = metaEventService.findAll(PageRequest.of(pageNum - 1, pageSize));
        
        Map<String, Object> response = new HashMap<>();
        response.put("events", events.getContent());
        response.put("pagination", Map.of(
            "page_num", pageNum,
            "page_size", pageSize,
            "total", events.getTotalElements(),
            "total_page", events.getTotalPages()
        ));
        
        return ApiResponse.success(response);
    }

    @GetMapping("/{id}")
    public ApiResponse<MetaEvent> findById(@PathVariable Long id) {
        MetaEvent event = metaEventService.findById(id);
        return ApiResponse.success(event);
    }

    @PostMapping
    public ApiResponse<MetaEvent> create(@RequestBody MetaEvent event) {
        MetaEvent createdEvent = metaEventService.create(event);
        return ApiResponse.success("Event created successfully", createdEvent);
    }

    @PostMapping("/{id}")
    public ApiResponse<MetaEvent> update(@PathVariable Long id, @RequestBody MetaEvent event) {
        MetaEvent updatedEvent = metaEventService.update(id, event);
        return ApiResponse.success("Event updated successfully", updatedEvent);
    }

    @PostMapping("/{id}/delete")
    public ApiResponse<Map<String, String>> delete(@PathVariable Long id) {
        metaEventService.delete(id);
        return ApiResponse.success("Event deleted successfully", 
            Map.of("message", "Event deleted successfully"));
    }
} 