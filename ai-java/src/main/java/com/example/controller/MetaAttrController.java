package com.example.controller;

import com.example.entity.MetaAttr;
import com.example.service.MetaAttrService;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api")
public class MetaAttrController {

    private final MetaAttrService metaAttrService;

    public MetaAttrController(MetaAttrService metaAttrService) {
        this.metaAttrService = metaAttrService;
    }

    @GetMapping("/attributes")
    public ResponseEntity<Page<MetaAttr>> findAll(
            @RequestParam(defaultValue = "1") int pageNum,
            @RequestParam(defaultValue = "10") int pageSize) {
        Page<MetaAttr> attributes = metaAttrService.findAll(PageRequest.of(pageNum - 1, pageSize));
        return ResponseEntity.ok(attributes);
    }

    @GetMapping("/attributes/{id}")
    public ResponseEntity<MetaAttr> findById(@PathVariable Long id) {
        MetaAttr attribute = metaAttrService.findById(id);
        return ResponseEntity.ok(attribute);
    }

    @PostMapping("/attributes")
    public ResponseEntity<MetaAttr> create(@RequestBody MetaAttr attribute) {
        MetaAttr createdAttribute = metaAttrService.create(attribute);
        return ResponseEntity.ok(createdAttribute);
    }

    @PostMapping("/attributes/{id}")
    public ResponseEntity<MetaAttr> update(@PathVariable Long id, @RequestBody MetaAttr attribute) {
        MetaAttr updatedAttribute = metaAttrService.update(id, attribute);
        return ResponseEntity.ok(updatedAttribute);
    }

    @PostMapping("/attributes/{id}/delete")
    public ResponseEntity<Map<String, String>> delete(@PathVariable Long id) {
        metaAttrService.delete(id);
        return ResponseEntity.ok(Map.of("message", "Attribute deleted successfully"));
    }

    @GetMapping("/attribute-types")
    public ResponseEntity<List<Map<String, String>>> getAttributeTypes() {
        return ResponseEntity.ok(metaAttrService.getAttributeTypes());
    }
} 