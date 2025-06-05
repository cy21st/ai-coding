package com.example.service;

import com.example.entity.MetaAttr;
import com.example.repository.MetaAttrRepository;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Service
public class MetaAttrService {

    private final MetaAttrRepository metaAttrRepository;
    private static final List<String> VALID_TYPES = Arrays.asList("string", "number", "bool", "time");

    public MetaAttrService(MetaAttrRepository metaAttrRepository) {
        this.metaAttrRepository = metaAttrRepository;
    }

    public Page<MetaAttr> findAll(Pageable pageable) {
        return metaAttrRepository.findByIsDeletedFalse(pageable);
    }

    public MetaAttr findById(Long id) {
        return metaAttrRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Attribute not found with id: " + id));
    }

    @Transactional
    public MetaAttr create(MetaAttr attr) {
        validateAttributeType(attr.getAttrType());
        return metaAttrRepository.save(attr);
    }

    @Transactional
    public MetaAttr update(Long id, MetaAttr attr) {
        validateAttributeType(attr.getAttrType());
        MetaAttr existingAttr = findById(id);
        existingAttr.setAttrName(attr.getAttrName());
        existingAttr.setAttrType(attr.getAttrType());
        existingAttr.setAttrDesc(attr.getAttrDesc());
        return metaAttrRepository.save(existingAttr);
    }

    @Transactional
    public void delete(Long id) {
        MetaAttr attr = findById(id);
        attr.setIsDeleted(true);
        metaAttrRepository.save(attr);
    }

    private void validateAttributeType(String type) {
        if (!VALID_TYPES.contains(type)) {
            throw new IllegalArgumentException("Invalid attribute type. Valid types are: " + VALID_TYPES);
        }
    }

    public List<Map<String, String>> getAttributeTypes() {
        return VALID_TYPES.stream()
                .map(type -> Map.of(
                        "value", type,
                        "label", getTypeLabel(type)
                ))
                .collect(Collectors.toList());
    }

    private String getTypeLabel(String type) {
        switch (type) {
            case "string": return "字符串";
            case "number": return "数字";
            case "bool": return "布尔值";
            case "time": return "日期";
            default: return type;
        }
    }
} 