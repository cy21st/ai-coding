package com.example.service;

import com.example.entity.MetaEvent;
import com.example.repository.MetaEventRepository;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class MetaEventService {

    private final MetaEventRepository metaEventRepository;

    public MetaEventService(MetaEventRepository metaEventRepository) {
        this.metaEventRepository = metaEventRepository;
    }

    public Page<MetaEvent> findAll(Pageable pageable) {
        return metaEventRepository.findByIsDeletedFalse(pageable);
    }

    public MetaEvent findById(Long id) {
        return metaEventRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Event not found with id: " + id));
    }

    @Transactional
    public MetaEvent create(MetaEvent event) {
        return metaEventRepository.save(event);
    }

    @Transactional
    public MetaEvent update(Long id, MetaEvent event) {
        MetaEvent existingEvent = findById(id);
        existingEvent.setEventName(event.getEventName());
        existingEvent.setEventDesc(event.getEventDesc());
        return metaEventRepository.save(existingEvent);
    }

    @Transactional
    public void delete(Long id) {
        MetaEvent event = findById(id);
        event.setIsDeleted(true);
        metaEventRepository.save(event);
    }
} 